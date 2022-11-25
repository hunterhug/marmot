package expert

// Package expert is use to parse content
import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
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

// FindPicture Find All picture. Must prefix with http(s)
func FindPicture(s string) []string {
	picList := make([]string, 0)
	re, _ := regexp.Compile(`src\s*=\s*["'](http[s]?:\/\/.*?\.(jpg|jpeg|png|gif))["']`)
	output := re.FindAllStringSubmatch(s, -1)
	for _, o := range output {
		picList = append(picList, o[1])
	}
	return picList
}
