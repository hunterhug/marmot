/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2020
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package expert

// Package expert is use to parse content
import (
	"regexp"
	"strings"

	"github.com/hunterhug/marmot/util/goquery"
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
	picList := make([]string, 0)
	re, _ := regexp.Compile(`src\s*=\s*["'](http[s]?:\/\/.*?\.(jpg|jpeg|png|gif))["']`)
	output := re.FindAllStringSubmatch(s, -1)
	for _, o := range output {
		picList = append(picList, o[1])
	}
	return picList
}
