package appstoreapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const asUserAgent = "AppStore/3.0 iOS/11.1.1 model/iPhone6,2 hwp/s5l8960x build/15B150 (6; dt:90)"
const asProxyURL = "http://176.9.112.168:5005"

var debug = false

// Story returns a Story by ID id.
func Story(id uint32, cc string, l string) (Page, error) {
	const errMsg = "[ERR] scraper.Story(%d,%s,%s): %v\n"
	const baseURL = "https://apps.apple.com/%s/story/id%d"
	uri, err := url.Parse(fmt.Sprintf(baseURL, cc, id))
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, nil
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, nil
	}
	query.Add("cc", cc)
	uri.RawQuery = query.Encode()

	storeFront := BuildStoreFront(cc, l)

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, nil
	}

	if resp.StatusCode != 200 {
		return Page{}, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, nil
	}

	page, err := ParsePage(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	return page, nil
}

// Suggestions returns suggestions by a keyword.
func Suggestions(keyword string, cc string, l string) []byte {
	const errMsg = "[ERR] scraper.Suggestions(%s,%s,%s): %v\n"
	const baseURL = "https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return nil
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return nil
	}
	query.Add("clientApplication", "Software")
	query.Add("caller", "com.apple.AppStore")
	query.Add("version", "1")
	query.Add("term", keyword)
	uri.RawQuery = query.Encode()

	storeFront := BuildStoreFront(cc, l)

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return nil
	}

	return body
}

// Genre returns App Store root page for Genre structure
func Genre(id uint32, cc string) (Page, error) {
	const errMsg = "[ERR] scraper.Genre(%d,%s): %v\n"
	const baseURL = "https://itunes.apple.com/%s/genre"
	uri, err := url.Parse(fmt.Sprintf(baseURL, cc))
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, err)
		return Page{}, err
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, err)
		return Page{}, err
	}
	query.Add("id", string(id))
	uri.RawQuery = query.Encode()

	storeFront := BuildStoreFront(cc, "")

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	proxyURL, err := url.Parse(asProxyURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, err)
		return Page{}, err
	}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	client = &http.Client{} // использовать ли прокси TODO
	if debug {
		log.Printf("[DBG] %s (%s): %s", cc, storeFront, req.URL.String())
	}

	resp, err := client.Do(req)
	if err != nil {
		if debug {
			fmt.Fprintf(os.Stderr, errMsg, id, cc, err)
		}
		return Page{}, err
	}

	if resp.StatusCode != 200 {
		return Page{}, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, err)
		return Page{}, err
	}

	page, err := ParsePage(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, err)
		return Page{}, err
	}

	return page, nil
}

// Grouping returns App Store root page for Grouping structure
func Grouping(id uint32, cc string, l string) (Page, error) {
	const errMsg = "[ERR] scraper.Grouping(%d,%s,%s): %v\n"
	const baseURL = "http://itunes.apple.com/WebObjects/MZStore.woa/wa/viewGrouping"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}
	query.Add("cc", cc)
	query.Add("id", string(id))
	uri.RawQuery = query.Encode()

	storeFront := BuildStoreFront(cc, l)

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	proxyURL, err := url.Parse(asProxyURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	client = &http.Client{} // TODO Использовать ли прокси
	if debug {
		log.Printf("[DBG] %s (%s): %s", cc, storeFront, req.URL.String())
	}

	resp, err := client.Do(req)
	if err != nil {
		if debug {
			fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		}
		return Page{}, err
	}

	if resp.StatusCode != 200 {
		return Page{}, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	page, err := ParsePage(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	return page, nil
}

// Room returns a Room by its ID.
func Room(id uint32, cc string, l string) (Page, error) {
	const errMsg = "[ERR] scraper.Room(%d,%s,%s): %v\n"
	const baseURL = "https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewRoom"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}
	query.Add("fcId", string(id))
	// query.Add("genreIdString", "6014")                           // TODO изучить
	// query.Add("mediaTypeString", "Mobile+Software+Applications") // TODO изучить
	uri.RawQuery = query.Encode()

	storeFront := BuildStoreFront(cc, l)

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	if resp.StatusCode != 200 {
		return Page{}, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	var page Page

	// TODO проверку на adamIds []string улучшить
	var adamIDsIsTypeOfStringArray bool
	err = json.Unmarshal(body, &page)
	if err != nil {
		adamIDsIsTypeOfStringArray = true
	}

	if adamIDsIsTypeOfStringArray {
		page, err = ParsePageAdamIDsString(body)
	} else {
		page, err = ParsePage(body)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	return page, nil
}

// App returns an application by ID id.
func App(id uint32, cc string, l string) (Page, error) {
	const errMsg = "[ERR] scraper.App(%d,%s,%s): %v\n"
	const baseURL = "https://apps.apple.com/%s/app/id%d"
	uri, err := url.Parse(fmt.Sprintf(baseURL, cc, id))
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}
	uri.RawQuery = query.Encode()

	storeFront := BuildStoreFront(cc, l)

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	if resp.StatusCode != 200 {
		return Page{}, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	page, err := ParsePage(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	return page, nil
}

// Bundle returns a bundle by ID id.
func Bundle(id uint32, cc string, l string) (Page, error) {
	const errMsg = "[ERR] scraper.Bundle(%d,%s,%s): %v\n"
	const baseURL = "https://apps.apple.com/%s/app-bundle/id%d"
	uri, err := url.Parse(fmt.Sprintf(baseURL, cc, id))
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}
	uri.RawQuery = query.Encode()

	storeFront := BuildStoreFront(cc, l)

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	if resp.StatusCode != 200 {
		return Page{}, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	page, err := ParsePage(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	return page, nil
}

// AppIDs returns application IDs by a keyword.
func AppIDs(keyword string, cc string, l string) []MetadataResponse {
	const errMsg = "[ERR] scraper.AsAppIDs(%s,%s,%s): %v\n"
	const baseURL = "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return []MetadataResponse{}
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return []MetadataResponse{}
	}
	query.Add("clientApplication", "Software")
	query.Add("caller", "com.apple.AppStore")
	query.Add("version", "1")
	query.Add("term", keyword)
	uri.RawQuery = query.Encode()

	storeFront := BuildStoreFront(cc, l)

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return []MetadataResponse{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return []MetadataResponse{}
	}

	return parseIDs(body)
}
