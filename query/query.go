/*
Copyright 2017 by GoSpider author. Email: gdccmcm14@live.com
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package query

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

// Find All picture. Must prefix http(s)://
func FindPicture(s string) []string {
	returnlist := []string{}
	re, _ := regexp.Compile(`src\s*=\s*["'](http[s]?:\/\/.*?\.(jpg|jpeg|png|gif))["']`)
	output := re.FindAllStringSubmatch(s, -1)
	for _, o := range output {
		returnlist = append(returnlist, o[1])
	}
	return returnlist
}
