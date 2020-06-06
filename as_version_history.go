package appstoreapi

import "time"

type VersionHistory []struct {
	ReleaseNotes  string    `json:"releaseNotes"`
	VersionString string    `json:"versionString"`
	ReleaseDate   time.Time `json:"releaseDate"`
}
