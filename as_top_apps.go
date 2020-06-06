package appstoreapi

type TopApps struct {
	Iphone TopAppsType `json:"iphone"`
}

type TopAppsType struct {
	Ids   []string `json:"ids"`
	Title string   `json:"title"`
}
