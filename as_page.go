package appstoreapi

// Page is a root App Store structure.
type Page struct {
	StorePlatformData map[string]StorePlatformData `json:"storePlatformData"`
	PageData          PageData                     `json:"pageData"`
	Properties        Properties                   `json:"properties"`
}
