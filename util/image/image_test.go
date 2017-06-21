/*
Copyright 2017 by GoSpider author.
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
package image

import (
	"testing"
)

func TestImage(t *testing.T) {

	// Scale a image file by cuting 100*100
	err := ThumbnailF2F("../data/image.png", "../data/image100-100.png", 100, 100)
	if err != nil {
		t.Error("Test ThumbnailF2F:" + err.Error())
	}

	// Scale a image file by cuting width:200 (Equal scaling)
	err = ScaleF2F("../data/image.png", "../data/image200.png", 200)
	if err != nil {
		t.Error("Test ScaleF2F:" + err.Error())
	}

	// File Real name
	filename, err := RealImageName("../data/image.png")
	if err != nil {
		t.Error("Test RealImageName:" + err.Error())
	} else {
		t.Log("Test RealImageName::real filename" + filename)
	}
}
