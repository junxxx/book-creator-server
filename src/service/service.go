package service

import (
	"github.com/junxxx/notion-book-creater/thirdparty"
)

// build the payload then call the notion api
func CreatePage(isbn string) error {
	item, err := thirdparty.GetBookDetail(isbn)
	if err != nil {
		return err
	}
	payload := thirdparty.NewPayload().
		AddTitle(item.Title).
		AddSubtitle(item.Subtitle).
		AddISBN(isbn).
		AddAuthor("Unknown")
	return thirdparty.CreatePage(payload.String())
}
