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
package myhbase

import (
	"fmt"
	"github.com/hunterhug/go-hbase"
	"testing"
	"time"
)

/*
before test,To do bellow

create_namespace 'hunterhug'
create 'hunterhug:info','colfamily',{SPLITALGO => 'HexStringSplit',NUMREGIONS => 40}
*/

func TestHbase(t1 *testing.T) {
	var (
		//格式化时间
		// Time format
		fs = "2006-01-02 15:04:05"

		//hbase配置
		//hbase config
		config = HbaseConfig{
			Zkport:   "2181", //zk
			Zkquorum: "192.168.11.73",
		}
		//创建一个客户端
		// new client
		clients = New(config)
	)

	//打开客户端
	//open connection
	clients.Open()

	//放主键
	//set a rowkey
	rowkey := "rowkey"

	//命名空间
	//set namespace
	hbasenamespace := "hunterhug"

	//表名
	//set tablename
	hbasetable := "info"

	//列族
	//set colfamily
	hbasefamily := "colfamily"

	//set col
	hbasecol := "go"
	hbasecol1 := "die"

	//放值，第一列
	// put a value to  hunterhug:info colfamily:go
	put := hbase.CreateNewPut([]byte(rowkey))
	put.AddStringValue(hbasefamily, hbasecol, "value")
	_, err := clients.Client.Put(hbasenamespace+":"+hbasetable, put)

	if err != nil {
		t1.Logf(err.Error())
	}

	//第二列
	// the same,put another value
	put.AddStringValue(hbasefamily, hbasecol1, "value1")
	_, err = clients.Client.Put(hbasenamespace+":"+hbasetable, put)

	if err != nil {
		t1.Logf(err.Error())
	}

	//获取一行
	//get value by rowkey
	result, err := clients.GetResult(hbasenamespace+":"+hbasetable, rowkey)
	if err != nil {
		t1.Logf(err.Error())
	}
	if v, exist := result.Columns[hbasefamily+":"+hbasecol]; exist {
		fmt.Printf("found row:%v,v:%v\n", rowkey, v.Value.String())
	}

	//扫描
	// create a scan object,so scan can be close because client is inside in it
	scan := clients.Client.Scan(hbasenamespace + ":" + hbasetable)

	//过滤两列，只拿
	// just scan one col
	err = scan.AddString(hbasefamily + ":" + hbasecol)

	if err != nil {
		panic(err)
	}

	// no!we want two col
	err = scan.AddString(hbasefamily + ":" + hbasecol1)
	if err != nil {
		panic(err)
	}

	//过滤时间戳，只拿时间戳在这段范围的合数据
	// and we are just want before 720h's data
	var t = time.Now().Add(-1 * 720 * time.Hour * 1)

	fmt.Printf("Get Data which record before %v\n", t.Format(fs))

	// Time filter
	scan.SetTimeRangeFrom(t)

	var v1, v2 string

	//循环处理值
	// get scan value
	scan.Map(func(r *hbase.ResultRow) {
		// if value exist,v1!
		if v, exist := r.Columns[hbasefamily+":"+hbasecol]; exist {
			v1 = v.Value.String()
		}
		if v, exist := r.Columns[hbasefamily+":"+hbasecol1]; exist {
			v2 = v.Value.String()
		}
		// output
		fmt.Printf("row:%s\t  %s:%s \t %s:%s \t", r.Row.String(), hbasecol, v1, hbasecol1, v2)
	})
}
