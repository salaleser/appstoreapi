package appstoreapi

import "time"

type Properties struct {
	RevNum    string    `json:"revNum"`
	Timestamp time.Time `json:"timestamp"`
}
