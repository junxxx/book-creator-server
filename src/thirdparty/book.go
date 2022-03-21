package thirdparty

import (
	"encoding/json"
	"net/http"
)

const url = `https://www.googleapis.com/books/v1/volumes`

type IndustryIdentifier struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

type VolumeInfo struct {
	Title               string               `json:"title"`
	Subtitle            string               `json:"subtitle"`
	Description         string               `json:"description"`
	IndustryIdentifiers []IndustryIdentifier `json:"industryIdentifiers"`
}

type Item struct {
	Kind       string     `json:"kind"`
	Id         string     `json:"id"`
	Etag       string     `json:"etag"`
	SelfLink   string     `json:"selfLink"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
}

type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}

// get the book detail information by ISBN
func GetBookDetail(isbn string) (VolumeInfo, error) {
	var detail VolumeInfo
	params := "?q=isbn:" + isbn
	url := url + params
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return detail, err
	}

	res, err := client.Do(req)
	if err != nil {
		return detail, err
	}
	defer res.Body.Close()

	var bookItem Response
	err = json.NewDecoder(res.Body).Decode(&bookItem)
	detail = bookItem.Items[0].VolumeInfo
	return detail, nil
}
