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
	"github.com/hunterhug/GoSpider/spider"
)

func main() {
	s, _ := spider.NewSpider(nil)
	// 打印spider对象
	fmt.Printf("%#v,%d\n", s, len(""))
	// 打印Ua
	fmt.Printf("%#v\n",spider.Ua)
	spider.UaInit()
	// 再打印
	fmt.Printf("%#v\n",spider.Ua)
	// 随机Ua
	fmt.Printf("%#v\n",spider.RandomUa())
	fmt.Printf("%#v\n",spider.RandomUa())
	fmt.Printf("%#v\n",spider.RandomUa())
	fmt.Printf("%#v\n",spider.RandomUa())
	fmt.Printf("%#v\n",spider.RandomUa())
	fmt.Printf("%#v\n",spider.RandomUa())
}
