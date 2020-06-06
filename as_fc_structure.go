package appstoreapi

type FcStructure struct {
	Version         int   `json:"version"`
	IsAuthenticated bool  `json:"isAuthenticated"`
	Model           Model `json:"model"`
	Meta            Meta  `json:"meta"`
}
