package appstoreapi

type Today struct {
	Metrics struct {
		RecoEligible  string `json:"reco_eligible"`
		RecoID        string `json:"reco_id"`
		RecoTimeStamp string `json:"reco_timeStamp"`
	} `json:"metrics"`
	Modules Modules `json:"modules"`
}
