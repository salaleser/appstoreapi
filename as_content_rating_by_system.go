package appstoreapi

type ContentRatingsBySystem struct {
	AppsApple System `json:"appsApple"`
}

type System struct {
	Name       string   `json:"name"`
	Value      int      `json:"value"`
	Rank       int      `json:"rank"`
	Advisories []string `json:"advisories"`
}
