package appstoreapi

import "time"

type DeveloperResponse struct {
	ID       int       `json:"id"`
	Body     string    `json:"body"`
	Modified time.Time `json:"modified"`
}
