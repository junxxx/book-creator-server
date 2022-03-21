package thirdparty

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	host = `https://api.notion.com`
)

var (
	secret      = os.Getenv("NOTION_SECRET")
	database_id = os.Getenv("NOTION_DATABASE_ID")
)

type Parent struct {
	Database_id string `json:"database_id"`
}

type Propertie struct {
	Title struct {
		Title []Text `json:"title"`
	} `json:"Title"`
	Subtitle RichText `json:"Subtitle"`
	Author   RichText `json:"Author"`
	ISBN     RichText `json:"ISBN"`
}

type Text struct {
	Text Content `json:"text"`
}

type Content struct {
	Content string `json:"content"`
}

type RichText struct {
	Rich_text []Text `json:"rich_text"`
}

type Status struct {
	Select_status struct {
		Name string `json:"name"`
	} `jsno:"status"`
}

type Payload struct {
	Parent     Parent    `json:"parent"`
	Properties Propertie `json:"properties"`
}

func NewPayload() *Payload {
	return &Payload{
		Parent: Parent{Database_id: database_id},
	}
}

// todo refactor
func (p *Payload) AddTitle(title string) *Payload {
	p.Properties.Title.Title = append(p.Properties.Title.Title, Text{Text: Content{title}})
	return p
}

func (p *Payload) AddSubtitle(subtitle string) *Payload {
	p.Properties.Subtitle.Rich_text = append(p.Properties.Subtitle.Rich_text, Text{Text: Content{subtitle}})
	return p
}

func (p *Payload) AddAuthor(author string) *Payload {
	p.Properties.Author.Rich_text = append(p.Properties.Author.Rich_text, Text{Text: Content{author}})
	return p
}

func (p *Payload) AddISBN(isbn string) *Payload {
	p.Properties.ISBN.Rich_text = append(p.Properties.ISBN.Rich_text, Text{Text: Content{isbn}})
	return p
}

func (p *Payload) String() string {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(p)
	return buf.String()
}

// create notoin page
func CreatePage(data string) error {
	url := host + "/v1/pages?"
	method := "POST"
	payload := strings.NewReader(data)
	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}

	// todo optimization
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Notion-Version", "2022-02-22")
	req.Header.Add("Authorization", "Bearer "+secret)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	r, err := ioutil.ReadAll(res.Body)
	log.Println(string(r))

	if res.StatusCode != http.StatusOK {
		return errors.New("notion api return " + strconv.Itoa(res.StatusCode))
	}
	log.Println(res.StatusCode)
	return nil
}
