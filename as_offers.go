package appstoreapi

type Offers []struct {
	ActionText     ActionText `json:"actionText"`
	Type           string     `json:"type"`
	PriceFormatted string     `json:"priceFormatted"`
	Price          float32    `json:"price"`
	BuyParams      string     `json:"buyParams"`
	Version        Version    `json:"version"`
	Assets         Assets     `json:"assets"`
}
