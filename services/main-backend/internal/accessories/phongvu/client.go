package phongvu

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

const (
	defaultTimeout        = 10 * time.Second
	maximumBodySize       = 5 << 20
	maximumCatalogDetails = 12
)

var nonDigitPattern = regexp.MustCompile(`[^0-9]`)

type Product struct {
	RetailerProductID string
	Name              string
	Brand             string
	Category          string
	ImageURL          string
	SourceURL         string
	Price             int64
	OriginalPrice     int64
	InStock           bool
}

type Client struct {
	httpClient *http.Client
	timeout    time.Duration
}

func NewClient(httpClient *http.Client, timeout time.Duration) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	if timeout <= 0 {
		timeout = defaultTimeout
	}
	return &Client{httpClient: httpClient, timeout: timeout}
}

func (client *Client) FetchCategory(ctx context.Context, sourceURL, category string) ([]Product, error) {
	body, err := client.fetchDocument(ctx, sourceURL)
	if err != nil {
		return nil, err
	}
	products, productErr := ParseCatalogHTML(strings.NewReader(string(body)), category)
	if productErr == nil {
		return products, nil
	}
	links, err := ParseCatalogLinks(strings.NewReader(string(body)))
	if err != nil {
		return nil, productErr
	}
	// ponytail: cap detail reads per category; paginate only when the demo catalog needs broader coverage.
	if len(links) > maximumCatalogDetails {
		links = links[:maximumCatalogDetails]
	}
	products = make([]Product, 0, len(links))
	for _, link := range links {
		detail, fetchErr := client.fetchDocument(ctx, link.URL)
		if fetchErr != nil {
			return nil, fetchErr
		}
		parsed, parseErr := ParseCatalogHTML(strings.NewReader(string(detail)), category)
		if parseErr != nil {
			return nil, parseErr
		}
		products = append(products, parsed...)
	}
	return products, nil
}

func (client *Client) fetchDocument(ctx context.Context, sourceURL string) ([]byte, error) {
	parsedURL, err := url.Parse(sourceURL)
	if err != nil || parsedURL.Scheme != "https" || (parsedURL.Hostname() != "phongvu.vn" && parsedURL.Hostname() != "www.phongvu.vn") {
		return nil, fmt.Errorf("source URL must use https://phongvu.vn")
	}
	requestContext, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()
	request, err := http.NewRequestWithContext(requestContext, http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create Phong Vu request: %w", err)
	}
	request.Header.Set("User-Agent", "SHOPWISE-AccessorySync/1.0")
	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("fetch Phong Vu category: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch Phong Vu category: status %d", response.StatusCode)
	}
	limited := io.LimitReader(response.Body, maximumBodySize+1)
	body, err := io.ReadAll(limited)
	if err != nil {
		return nil, fmt.Errorf("read Phong Vu category: %w", err)
	}
	if len(body) > maximumBodySize {
		return nil, fmt.Errorf("Phong Vu response exceeds 5 MB")
	}
	return body, nil
}

type CatalogLink struct {
	Name string
	URL  string
}

func ParseCatalogLinks(reader io.Reader) ([]CatalogLink, error) {
	document, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("parse catalog HTML: %w", err)
	}
	links := make([]CatalogLink, 0)
	var parseErrors []error
	var visit func(*html.Node)
	visit = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "script" && attribute(node, "type") == "application/ld+json" && node.FirstChild != nil {
			var value any
			if err := json.Unmarshal([]byte(node.FirstChild.Data), &value); err != nil {
				parseErrors = append(parseErrors, err)
			} else {
				collectCatalogLinks(value, &links)
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			visit(child)
		}
	}
	visit(document)
	if len(links) == 0 {
		return nil, fmt.Errorf("parse Phong Vu catalog links: %w", errors.Join(parseErrors...))
	}
	return links, nil
}

func collectCatalogLinks(value any, links *[]CatalogLink) {
	switch typed := value.(type) {
	case []any:
		for _, item := range typed {
			collectCatalogLinks(item, links)
		}
	case map[string]any:
		if schemaType(typed["@type"]) == "ListItem" {
			if item, ok := typed["item"].(map[string]any); ok && schemaType(item["@type"]) != "Product" {
				link := CatalogLink{Name: stringValue(item["name"]), URL: absolutePhongVuURL(stringValue(item["url"]))}
				if link.URL != "" {
					*links = append(*links, link)
				}
				return
			}
		}
		for _, item := range typed {
			collectCatalogLinks(item, links)
		}
	}
}

func ParseCatalogHTML(reader io.Reader, category string) ([]Product, error) {
	document, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("parse catalog HTML: %w", err)
	}
	var products []Product
	var parseErrors []error
	var visit func(*html.Node)
	visit = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "script" && node.FirstChild != nil && (attribute(node, "type") == "application/ld+json" || attribute(node, "id") == "__NEXT_DATA__") {
			decoder := json.NewDecoder(strings.NewReader(node.FirstChild.Data))
			decoder.UseNumber()
			var value any
			if err := decoder.Decode(&value); err != nil {
				parseErrors = append(parseErrors, err)
			} else if attribute(node, "id") == "__NEXT_DATA__" {
				collectNextProducts(value, category, &products)
			} else {
				collectProducts(value, category, &products)
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			visit(child)
		}
	}
	visit(document)
	if len(products) == 0 {
		if len(parseErrors) > 0 {
			return nil, fmt.Errorf("parse Phong Vu structured data: %w", errors.Join(parseErrors...))
		}
		return nil, fmt.Errorf("Phong Vu catalog did not contain structured products")
	}
	deduplicated := make([]Product, 0, len(products))
	seen := make(map[string]struct{}, len(products))
	for _, product := range products {
		if product.RetailerProductID == "" || product.Price <= 0 {
			continue
		}
		if _, exists := seen[product.RetailerProductID]; exists {
			continue
		}
		seen[product.RetailerProductID] = struct{}{}
		deduplicated = append(deduplicated, product)
	}
	if len(deduplicated) == 0 {
		return nil, fmt.Errorf("Phong Vu catalog contained no valid offers")
	}
	return deduplicated, nil
}

func collectNextProducts(value any, category string, products *[]Product) {
	root, _ := value.(map[string]any)
	props, _ := root["props"].(map[string]any)
	pageProps, _ := props["pageProps"].(map[string]any)
	items, _ := pageProps["serverProducts"].([]any)
	for _, item := range items {
		product, _ := item.(map[string]any)
		price, _ := product["price"].(map[string]any)
		link, _ := product["link"].(map[string]any)
		linkAs, _ := link["as"].(map[string]any)
		brand, _ := product["brand"].(map[string]any)
		*products = append(*products, Product{
			RetailerProductID: stringValue(product["sku"]), Name: stringValue(product["name"]),
			Brand: stringValue(brand["name"]), Category: category, ImageURL: stringValue(product["imageUrl"]),
			SourceURL: absolutePhongVuURL(stringValue(linkAs["pathname"])), Price: moneyValue(price["latestPrice"]),
			OriginalPrice: moneyValue(price["supplierRetailPrice"]), InStock: moneyValue(product["stockQuantity"]) > 0,
		})
	}
}

func collectProducts(value any, category string, products *[]Product) {
	switch typed := value.(type) {
	case []any:
		for _, item := range typed {
			collectProducts(item, category, products)
		}
	case map[string]any:
		if schemaType(typed["@type"]) == "Product" {
			*products = append(*products, decodeProduct(typed, category))
			return
		}
		for _, item := range typed {
			collectProducts(item, category, products)
		}
	}
}

func decodeProduct(value map[string]any, category string) Product {
	offer, _ := value["offers"].(map[string]any)
	brand := stringValue(value["brand"])
	if brandMap, ok := value["brand"].(map[string]any); ok {
		brand = stringValue(brandMap["name"])
	}
	return Product{
		RetailerProductID: firstNonEmpty(stringValue(value["sku"]), stringValue(value["productID"])),
		Name:              stringValue(value["name"]), Brand: brand, Category: category,
		ImageURL: imageValue(value["image"]), SourceURL: absolutePhongVuURL(stringValue(offer["url"])),
		Price: moneyValue(offer["price"]), OriginalPrice: moneyValue(offer["highPrice"]),
		InStock: strings.Contains(strings.ToLower(stringValue(offer["availability"])), "instock"),
	}
}

func moneyValue(value any) int64 {
	switch typed := value.(type) {
	case json.Number:
		integer, _ := strconv.ParseInt(strings.Split(typed.String(), ".")[0], 10, 64)
		return integer
	case float64:
		return int64(typed)
	case string:
		digits := nonDigitPattern.ReplaceAllString(typed, "")
		integer, _ := strconv.ParseInt(digits, 10, 64)
		return integer
	default:
		return 0
	}
}

func schemaType(value any) string {
	if values, ok := value.([]any); ok && len(values) > 0 {
		return stringValue(values[0])
	}
	return stringValue(value)
}

func stringValue(value any) string {
	text, _ := value.(string)
	return strings.TrimSpace(text)
}

func imageValue(value any) string {
	if values, ok := value.([]any); ok && len(values) > 0 {
		return stringValue(values[0])
	}
	return stringValue(value)
}

func absolutePhongVuURL(value string) string {
	parsed, err := url.Parse(value)
	if err != nil {
		return ""
	}
	base, _ := url.Parse("https://phongvu.vn")
	resolved := base.ResolveReference(parsed)
	if resolved.Hostname() != "phongvu.vn" && resolved.Hostname() != "www.phongvu.vn" {
		return ""
	}
	return resolved.String()
}

func attribute(node *html.Node, name string) string {
	for _, item := range node.Attr {
		if item.Key == name {
			return item.Val
		}
	}
	return ""
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}
