package phongvu

import (
	"context"
	"fmt"
	"time"

	ordersdomain "shopwise/retail/internal/orders/domain"
)

type OfferRefresher struct {
	client  *Client
	enabled bool
	now     func() time.Time
}

func NewOfferRefresher(client *Client, enabled bool, now func() time.Time) *OfferRefresher {
	return &OfferRefresher{client: client, enabled: enabled, now: now}
}

func (refresher *OfferRefresher) RefreshRetailerOffer(
	ctx context.Context,
	quote ordersdomain.RetailerOfferQuote,
) (ordersdomain.RetailerOfferQuote, error) {
	if !refresher.enabled {
		return ordersdomain.RetailerOfferQuote{}, fmt.Errorf("Phong Vu connector is disabled")
	}
	products, err := refresher.client.FetchCategory(ctx, quote.SourceURL, quote.Category)
	if err != nil {
		return ordersdomain.RetailerOfferQuote{}, err
	}
	for _, product := range products {
		if product.RetailerProductID != quote.RetailerProductID {
			continue
		}
		quote.UnitPrice = product.Price
		quote.OriginalPrice = product.OriginalPrice
		quote.InStock = product.InStock
		quote.SourceURL = product.SourceURL
		quote.FetchedAt = refresher.now().UTC()
		return quote, nil
	}
	return ordersdomain.RetailerOfferQuote{}, fmt.Errorf("Phong Vu offer %s was not found", quote.RetailerProductID)
}
