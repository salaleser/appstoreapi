package appstoreapi

// StorePlatformData contains results and can be: product-dv,
// native-search-lockup, lockup, editorial-item-product.
type StorePlatformData struct {
	Results         map[string]Result `json:"results"`
	Version         int               `json:"version"`
	IsAuthenticated bool              `json:"isAuthenticated"`
	Meta            Meta              `json:"meta"`
}
