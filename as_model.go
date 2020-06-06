package appstoreapi

type Model struct {
	AdamID             string  `json:"adamId"`
	Artwork            Artwork `json:"artwork,omitempty"`
	Children           []Model `json:"children,omitempty"`
	Content            []Link  `json:"content,omitempty"`
	DesignBadge        string  `json:"designBadge,omitempty"`
	DesignTag          string  `json:"designTag,omitempty"`
	DisplayStyle       string  `json:"displayStyle,omitempty"`
	DoNotFilter        bool    `json:"doNotFilter"`
	FcKind             string  `json:"fcKind"`
	Link               Link    `json:"link,omitempty"`
	Links              []Link  `json:"links,omitempty"`
	Name               string  `json:"name,omitempty"`
	SeeAllURL          string  `json:"seeAllUrl,omitempty"`
	ShouldUseGradients bool    `json:"shouldUseGradients"`
	SuppressTagline    string  `json:"suppressTagline,omitempty"`
	Tagline            string  `json:"tagline,omitempty"`
	Title              string  `json:"title,omitempty"`
}
