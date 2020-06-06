package appstoreapi

type Bubbles []struct {
	Results []struct {
		Type   int    `json:"type"`
		ID     string `json:"id"`
		Entity string `json:"entity"`
	} `json:"results"`
	Name       string `json:"name"`
	TotalCount int    `json:"totalCount"`
}
