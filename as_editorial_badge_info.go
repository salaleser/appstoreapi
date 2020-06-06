package appstoreapi

// EditorialBadgeInfo contains editorial badge info.
// EditorialBadgeType can be: editorialPriority...
type EditorialBadgeInfo struct {
	EditorialBadgeType string `json:"editorialBadgeType"`
	NameForDisplay     string `json:"nameForDisplay"`
}
