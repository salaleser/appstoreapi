package appstoreapi

type ActionText struct {
	Short       string `json:"short"`
	Medium      string `json:"medium"`
	Long        string `json:"long"`
	Downloaded  string `json:"downloaded"`
	Downloading string `json:"downloading"`
}
