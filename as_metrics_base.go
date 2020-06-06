package appstoreapi

type MetricsBase struct {
	PageType              string `json:"pageType"`
	PageID                string `json:"pageId"`
	PageDetails           string `json:"pageDetails"`
	Page                  string `json:"page"`
	ServerInstance        string `json:"serverInstance"`
	StoreFrontHeader      string `json:"storeFrontHeader"`
	Language              string `json:"language"`
	PlatformID            string `json:"platformId"`
	PlatformName          string `json:"platformName"`
	StoreFront            string `json:"storeFront"`
	EnvironmentDataCenter string `json:"environmentDataCenter"`
}
