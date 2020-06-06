package appstoreapi

type ChartPosition struct {
	Position  int    `json:"position"`
	GenreName string `json:"genreName"`
	ChartURL  string `json:"chartUrl"`
}
