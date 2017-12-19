// Pacakge expert is use to parse content
package expert

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery" // please include by yourself
)

func QueryBytes(content []byte) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	return doc, err
}

func QueryString(content string) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	return doc, err
}

// Find All picture. Must prefix with http(s)
func FindPicture(s string) []string {
	returnlist := []string{}
	re, _ := regexp.Compile(`src\s*=\s*["'](http[s]?:\/\/.*?\.(jpg|jpeg|png|gif))["']`)
	output := re.FindAllStringSubmatch(s, -1)
	for _, o := range output {
		returnlist = append(returnlist, o[1])
	}
	return returnlist
}
