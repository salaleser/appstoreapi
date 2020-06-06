package appstoreapi

// MetadataResponse is a vitalina's Application's metadata structure.
type MetadataResponse struct { // TODO add more fields
	Title       string
	Link        string
	AppID       string
	ArtistName  string
	Rating      float32
	ReleaseDate string
	Subtitle    string
	Description string
	Screenshot1 string // TODO add array
	Logo        string
}
