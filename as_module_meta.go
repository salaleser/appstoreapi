package appstoreapi

type ModuleMeta struct {
	DataRecoDataSetID  string `json:"data.reco.dataSetId"`
	RecoAdditionalData string `json:"reco_additionalData"`
	RecoEligible       string `json:"reco_eligible"`
	RecoID             string `json:"reco_id"`
	RecoTimeStamp      string `json:"reco_timeStamp"`
}
