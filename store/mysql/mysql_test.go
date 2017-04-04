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
package mysql

import (
	"testing"
)

/*

CREATE TABLE IF NOT EXISTS `51job_keyword` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `keyword` varchar(255) NOT NULL DEFAULT '',
  `address` varchar(255) NOT NULL DEFAULT '',
  `kind` varchar(255) NOT NULL DEFAULT '',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `time51` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='关键字表';

*/

func TestMysql(t *testing.T) {

	// mysql config
	config := MysqlConfig{
		Username: "root",
		Password: "smart2016",
		Ip:       "127.0.0.1",
		Port:     "3306",
		Dbname:   "51job",
	}
	e := config.CreateDb()
	if e != nil {
		t.Error(e.Error())
	}
	//e = config.DeleteDb()
	//if e != nil {
	//	t.Error(e.Error())
	//}
	// a new db connection
	db := New(config)

	// open connection
	db.Open(2000, 1000)

	// create sql
	sql := `
  CREATE TABLE IF NOT EXISTS 51job.51job_keyword (
  id int(11) NOT NULL AUTO_INCREMENT,
  keyword varchar(255) NOT NULL DEFAULT '',
  address varchar(255) NOT NULL DEFAULT '',
  kind varchar(255) NOT NULL DEFAULT '',
  created datetime DEFAULT NULL,
  updated datetime DEFAULT NULL,
  time51 int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='关键字表';`

	// create
	inum, err := db.Create(sql)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("create number:%d\n", inum)
	}

	// insert sql
	//'1', '教师', '潮州', '0', '2016-05-27 00:00:00', '2016-05-27 00:00:00', '204'
	sql = "INSERT INTO `51job_keyword`(`keyword`,`address`,`kind`) values(?,?,?)"

	// insert
	num, err := db.Insert(sql, "PHP", "潮州", 0)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("insert number:%d\n", num)
	}

	// select sql
	sql = "SELECT * FROM 51job_keyword where address=? and kind=? limit ?;"

	// select
	result, err := db.Select(sql, "潮州", 0, 6)
	if err != nil {
		t.Error(err.Error())
	} else {
		// values
		for row, v := range result {
			t.Logf("%v:%#v\n", row, v)
		}
	}
}
