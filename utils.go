package appstoreapi

import (
	"fmt"
	"strings"
)

func BuildStoreFront(cc string, l string) string {
	p, ok := PlatformIDs["iphone"]
	if !ok {
		p = 29
	}

	sf, ok := StoreFronts[strings.ToUpper(cc)]
	if !ok {
		return ""
	}

	asLanguageCode, ok := Languages[l]
	if !ok {
		return fmt.Sprintf("%d,%d t:apps3", sf, p)
	}

	return fmt.Sprintf("%d-%d,%d t:apps3", sf, asLanguageCode, p)
}
