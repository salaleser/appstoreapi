package appstoreapi

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func parseIDs(body []byte) ([]MetadataResponse, error) {
	const errMsg = "[ERR] scraper.parseAsIDs(%s...): %v\n"
	var data Page
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Fprintf(os.Stderr, errMsg, body[:10], err)
		return []MetadataResponse{}, err
	}

	metadatas := make([]MetadataResponse, 0)
	for _, result := range data.StorePlatformData["native-search-lockup"].Results {
		if result.Kind != "iosSoftware" {
			continue
		}

		metadata := MetadataResponse{
			Title:  result.Name,
			AppID:  result.ID,
			Rating: result.UserRating.Value,
		}

		metadatas = append(metadatas, metadata)
	}

	return metadatas, nil
}

// ParsePage parses App Store root page and returns structure.
func ParsePage(body []byte) (Page, error) {
	var page Page
	err := json.Unmarshal(body, &page)
	if err != nil {
		return Page{}, err
	}

	return page, nil
}

// ParsePageAdamIDsString converts "adamIds" []string to []int.
func ParsePageAdamIDsString(body []byte) (Page, error) {
	fmt.Fprintf(os.Stderr, "Room \"adamIds\" of type []string.\n")

	var page Page

	var page2 Page2
	err := json.Unmarshal(body, &page2)
	if err != nil {
		return Page{}, err
	}

	page2adamIDs := make([]int, 0)
	for _, adamID := range page2.PageData.AdamIDs {
		id, _ := strconv.Atoi(adamID)
		page2adamIDs = append(page2adamIDs, id)
	}
	page.StorePlatformData = page2.StorePlatformData
	page.PageData.AdamID = page2.PageData.AdamID
	page.PageData.ComponentName = page2.PageData.ComponentName
	page.PageData.Metrics = page2.PageData.Metrics
	page.PageData.Mt = page2.PageData.Mt
	page.PageData.PageTitle = page2.PageData.PageTitle
	page.PageData.FcKind = page2.PageData.FcKind
	page.PageData.MetricsBase = page2.PageData.MetricsBase
	page.PageData.AdamIDs = page2adamIDs
	page.Properties = page2.Properties

	return page, nil
}
