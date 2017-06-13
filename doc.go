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

package main

import (
	"fmt"
	_ "github.com/hunterhug/GoSpider/query"
	_ "github.com/hunterhug/GoSpider/spider"
	_ "github.com/hunterhug/GoSpider/store"
	_ "github.com/hunterhug/GoSpider/store/myetcd"
	_ "github.com/hunterhug/GoSpider/store/myredis"
	_ "github.com/hunterhug/GoSpider/store/mysql"
	_ "github.com/hunterhug/GoSpider/util"
)

func main() {
	fmt.Println("Hello GoSpider")
}
