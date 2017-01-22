package query

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func QueryBytes(content []byte) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	return doc, err
}

func QueryString(content string) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	return doc, err
}
