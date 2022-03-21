package thirdparty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestDetail(t *testing.T) {
	// isbn := "9787532772322"
	// GetBookDetail(isbn)
}

func TestCreatePage(t *testing.T) {
	payload := `
	{"parent":{"database_id":"8723c4cd560b49fa9328f674dd2cf5ad"},"properties":{"Title":{"title":[{"text":{"content":"test from api"}}]},"Subtitle":{"rich_text":[{"text":{"content":"A dark green leafy vegetable"}}]},"Author":{"rich_text":[{"text":{"content":"A dark green leafy "}}]},"Status":{"select":{"name":"Want to read"}},"ISBN":{"rich_text":[{"text":{"content":"9787532772322 "}}]}}}
	`

	CreatePage(payload)
}

func TestPayload(t *testing.T) {
	payload := `
	{"parent":{"database_id":"8723c4cd560b49fa9328f674dd2cf5ad"},"properties":{"Title":{"title":[{"text":{"content":"test from api"}}]},"Subtitle":{"rich_text":[{"text":{"content":"A dark green leafy vegetable"}}]},"Author":{"rich_text":[{"text":{"content":"A dark green leafy "}}]},"Status":{"select":{"name":"Want to read"}},"ISBN":{"rich_text":[{"text":{"content":"9787532772322 "}}]}}}
	`

	var p Payload
	err := json.NewDecoder(strings.NewReader(payload)).Decode(&p)
	fmt.Println(err)
	fmt.Printf("%T, %#v", p, p)

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(p)
	fmt.Println(err, buf.String())

}
