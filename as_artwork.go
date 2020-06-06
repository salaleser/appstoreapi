package appstoreapi

// Artwork contains image info.
// Can be: subscriptionHero, brandLogo, originalFlowcaseBrick, storeFlowcase,
// bannerUber, dayCard, mediaCard, generalCard, emailFeature.
type Artwork struct {
	Width                int    `json:"width"`
	URL                  string `json:"url"`
	Height               int    `json:"height"`
	TextColor3           string `json:"textColor3"`
	TextColor2           string `json:"textColor2"`
	TextColor4           string `json:"textColor4"`
	HasAlpha             bool   `json:"hasAlpha"`
	TextColor1           string `json:"textColor1"`
	BgColor              string `json:"bgColor"`
	HasP3                bool   `json:"hasP3"`
	SupportsLayeredImage bool   `json:"supportsLayeredImage"`
}
