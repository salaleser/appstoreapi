package appstoreapi

type Modules []struct {
	Date     string     `json:"date"`
	Meta     ModuleMeta `json:"meta"`
	Contents []string   `json:"contents"`
}
