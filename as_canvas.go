package appstoreapi

type Canvas []struct {
	DisplayType            string  `json:"displayType"`
	EditorialCopy          string  `json:"editorialCopy,omitempty"`
	Artwork                Artwork `json:"artwork,omitempty"`
	InlineImageDisplayType string  `json:"inlineImageDisplayType,omitempty"`
	ContentID              string  `json:"contentId,omitempty"`
	AppLockupSize          string  `json:"appLockupSize,omitempty"`
	AppLockupVideo         string  `json:"appLockupVideo,omitempty"`
}
