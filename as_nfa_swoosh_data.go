package appstoreapi

type NfaSwooshData struct {
	LockupCountMinimumDesktop int    `json:"lockupCountMinimumDesktop"`
	LockupCountMinimumIos     int    `json:"lockupCountMinimumIos"`
	SeeAllURL                 string `json:"seeAllUrl"`
	SeeAllURLThreshold        int    `json:"seeAllUrlThreshold"`
	LookupURLData             struct {
		BaseURL         string `json:"baseUrl"`
		TimeoutInMillis string `json:"timeoutInMillis"`
	} `json:"lookupUrlData"`
}
