package phongvu

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

const catalogFixture = `<!doctype html><html><head>
<script type="application/ld+json">{
  "@context":"https://schema.org","@type":"ItemList","itemListElement":[
    {"@type":"ListItem","item":{"@type":"Product","sku":"PV-MOUSE-1","name":"Chuột Logitech G304","brand":{"name":"Logitech"},"image":"https://phongvu.vn/media/g304.jpg","offers":{"price":"890.000 ₫","priceCurrency":"VND","availability":"https://schema.org/InStock","url":"/p/chuot-logitech-g304"}}},
    {"@type":"ListItem","item":{"@type":"Product","sku":"PV-MOUSE-1","name":"duplicate","offers":{"price":890000,"url":"/p/duplicate"}}}
  ]
}</script></head></html>`

func TestParseCatalogHTMLReadsVNDOfferAndDeduplicates(t *testing.T) {
	t.Parallel()

	products, err := ParseCatalogHTML(strings.NewReader(catalogFixture), "mouse")
	if err != nil {
		t.Fatalf("ParseCatalogHTML() error = %v", err)
	}
	if len(products) != 1 {
		t.Fatalf("len(products) = %d, want 1", len(products))
	}
	product := products[0]
	if product.RetailerProductID != "PV-MOUSE-1" || product.Price != 890_000 || !product.InStock {
		t.Fatalf("product = %#v", product)
	}
	if product.SourceURL != "https://phongvu.vn/p/chuot-logitech-g304" || product.Category != "mouse" {
		t.Fatalf("source/category = %s/%s", product.SourceURL, product.Category)
	}
}

func TestParseCatalogHTMLReadsCurrentNextDataCatalog(t *testing.T) {
	t.Parallel()

	fixture := `<!doctype html><html><body><script id="__NEXT_DATA__" type="application/json">{
  "props":{"pageProps":{"serverProducts":[{
    "sku":"260404921","name":"Mouse Inphic W1S","imageUrl":"https://phongvu.vn/media/w1s.jpg",
    "price":{"latestPrice":149000,"supplierRetailPrice":259000},"stockQuantity":4,
    "link":{"as":{"pathname":"/chuot-inphic-w1s--s260404921"}},"brand":{"name":"Inphic"}
  }]}}
}</script></body></html>`

	products, err := ParseCatalogHTML(strings.NewReader(fixture), "mouse")
	if err != nil {
		t.Fatalf("ParseCatalogHTML() error = %v", err)
	}
	if len(products) != 1 {
		t.Fatalf("len(products) = %d, want 1", len(products))
	}
	product := products[0]
	if product.RetailerProductID != "260404921" || product.Price != 149_000 || product.OriginalPrice != 259_000 || !product.InStock {
		t.Fatalf("product = %#v", product)
	}
	if product.SourceURL != "https://phongvu.vn/chuot-inphic-w1s--s260404921" {
		t.Fatalf("SourceURL = %q", product.SourceURL)
	}
}

func TestParseCatalogHTMLRejectsMalformedCatalog(t *testing.T) {
	t.Parallel()

	_, err := ParseCatalogHTML(strings.NewReader(`<html><script type="application/ld+json">{broken</script></html>`), "mouse")
	if err == nil {
		t.Fatal("ParseCatalogHTML() error = nil, want malformed data error")
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (function roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return function(request)
}

func TestClientHonorsRequestTimeout(t *testing.T) {
	t.Parallel()

	client := NewClient(&http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		<-request.Context().Done()
		return nil, request.Context().Err()
	})}, 5*time.Millisecond)
	_, err := client.FetchCategory(context.Background(), "https://phongvu.vn/c/chuot-co-day", "mouse")
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("FetchCategory() error = %v, want deadline exceeded", err)
	}
}

func TestClientRejectsNonPhongVuHost(t *testing.T) {
	t.Parallel()

	client := NewClient(&http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(catalogFixture))}, nil
	})}, time.Second)
	_, err := client.FetchCategory(context.Background(), "https://example.com/catalog", "mouse")
	if err == nil {
		t.Fatal("FetchCategory() error = nil, want host validation error")
	}
}

func TestClientFollowsPublicCatalogLinksToProductDetails(t *testing.T) {
	t.Parallel()

	catalog := `<script type="application/ld+json">{"@type":"ItemList","itemListElement":[{"@type":"ListItem","item":{"name":"Mouse","url":"https://phongvu.vn/mouse--p1"}}]}</script>`
	detail := `<script type="application/ld+json">{"@type":"Product","sku":"PV-DETAIL-1","name":"Mouse","offers":{"price":"239000","availability":"https://schema.org/InStock","url":"https://phongvu.vn/mouse--p1"}}</script>`
	client := NewClient(&http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		body := catalog
		if request.URL.Path == "/mouse--p1" {
			body = detail
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(body))}, nil
	})}, time.Second)

	products, err := client.FetchCategory(context.Background(), "https://phongvu.vn/c/mouse", "mouse")
	if err != nil {
		t.Fatalf("FetchCategory() error = %v", err)
	}
	if len(products) != 1 || products[0].RetailerProductID != "PV-DETAIL-1" || products[0].Price != 239_000 {
		t.Fatalf("products = %#v", products)
	}
}
