package appstoreapi

// Page2 это костыль для конвертации appIds []string в []int.
type Page2 struct {
	StorePlatformData map[string]StorePlatformData `json:"storePlatformData"`
	PageData          PageData2                    `json:"pageData"`
	Properties        Properties                   `json:"properties"`
}
