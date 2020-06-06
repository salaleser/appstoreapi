package appstoreapi

type Genres []struct {
	GenreID       string `json:"genreId"`
	Name          string `json:"name"`
	URL           string `json:"url"`
	MediaType     string `json:"mediaType"`
	ParentGenreID string `json:"parentGenreId"`
}
