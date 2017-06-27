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
package mydb

import (
	"fmt"
	"testing"
	"time"
)

type JobInfo struct {
	Id          string    `xorm:"varchar(100) pk"`
	Handle      string    `xorm:"varchar(100) notnull"`
	InstanceId  string    `xorm:"varchar(100) notnull"`
	ContainerId string    `xorm:"varchar(100) notnull"`
	Data        string    `xorm:"text notnull"`
	Created     time.Time `xorm:"DateTime notnull"`
	UpdateTime  time.Time `xorm:"DateTime updated"`
}

func TestNewDb(t *testing.T) {

	config := MyDbConfig{
		DriverName: PG,
		DbConfig: DbConfig{
			Host: "127.0.0.1",
			User: "jinhan",
			Pass: "6833066",
			Name: "jinhan",
		},
		MaxOpenConns: 1,
		MaxIdleConns: 1,
		Debug:        true,
	}

	db, err := NewDb(config)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		err := db.Ping()
		if err != nil {
			fmt.Println("d" + err.Error())
		}
	}

	x, err := db.Client.DBMetas()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%#v", x)
	}

	ok, err := db.Client.IsTableExist(JobInfo{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ok)
	}

	db.DropTables(JobInfo{})
	err = db.Client.CreateTables(JobInfo{})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMyDb_Insert(t *testing.T) {
	config := MyDbConfig{
		DriverName: PG,
		DbConfig: DbConfig{
			Host: "127.0.0.1",
			User: "jinhan",
			Pass: "6833066",
			Name: "jinhan",
		},
		MaxOpenConns: 1,
		MaxIdleConns: 1,
		Debug:        true,
	}

	db, err := NewDb(config)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		err := db.Ping()
		if err != nil {
			fmt.Println("d" + err.Error())
		}
	}

	job := JobInfo{
		Id:          "aaddsdtttaa-bbbbbbbbddd",
		ContainerId: "a",
		InstanceId:  "aaaaaa",
		Handle:      "bbbbbbb",
		Created:     time.Now(),
		Data:        "asfasssssssssssssssssssssss",
	}
	job1 := JobInfo{
		Id:          "aaaddddaaa-bbbbbbbb",
		ContainerId: "sssssssss",
		InstanceId:  "aaaaaa",
		Handle:      "bbbbbbb",
		Created:     time.Now(),
		Data:        "asfasssssssssssssssssssssss",
	}
	fmt.Println(job, job1)
	/*
	   INSERT INTO table (id, field, field2) VALUES (1, A, X), (2, B, Y), (3, C, Z)
	   ON DUPLICATE KEY UPDATE field=VALUES(Col1), field2=VALUES(Col2);
	*/
	_ = `INSERT INTO job_info(id,handle,instance_id,container_id,data,created,update_time) VALUES(?,?,?,?,?,?,?)
	ON DUPLICATE KEY UPDATE update_time=?;`

	sql := `INSERT INTO job_info(id,handle,instance_id,container_id,data,created,update_time) VALUES(?,?,?,?,?,?,?)
	ON CONFLICT (id) DO UPDATE SET update_time=?;
	`
	_, e := db.Exec(sql, job.Id, job.Handle, job.InstanceId, job.ContainerId, job.Data, job.Created, "2017-06-16 17:29:35.974709", "2017-06-16 17:29:35.974709")
	if e != nil {
		fmt.Println(e.Error())
	}
	/*
		INSERT INTO the_table (id, column_1, column_2)
		VALUES (1, 'A', 'X'), (2, 'B', 'Y'), (3, 'C', 'Z')
		ON CONFLICT (id) DO UPDATE
		  SET column_1 = excluded.column_1,
		      column_2 = excluded.column_2;
	*/

}
