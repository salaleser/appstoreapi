package appstoreapi

type Link struct {
	ContentID string `json:"contentId"`
	KindID    string `json:"kindId"`
	KindIDs   []int  `json:"kindIds"`
	Label     string `json:"label"`
	Target    string `json:"target"`
	Type      string `json:"type"`
	URL       string `json:"url"`
}
