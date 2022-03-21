package service

import (
	"fmt"
	"testing"
)

func TestCreatPage(t *testing.T) {
	isbn := "9787532772322"
	err := CreatePage(isbn)
	fmt.Println(err)
}
