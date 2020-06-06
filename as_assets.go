package appstoreapi

type Assets []struct {
	Flavor string `json:"flavor"`
	Size   int    `json:"size"`
}
