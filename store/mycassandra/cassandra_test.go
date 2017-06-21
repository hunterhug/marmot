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
package mycassandra

import (
	"github.com/gocql/gocql"
	"testing"
)

/*
测试之前先填充一下cassandra语句
Before test:

create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
create index on example.tweet(timeline);
*/
func TestCdb(t *testing.T) {
	//先构造一个字符串数组连接
	// cassandra host
	host := []string{"192.168.11.74"}

	//指定cassandra keyspace，类似于mysql中的db
	// cassandra keyspace like mysql db
	keyword := "example"

	//连接
	// connect a cdb
	cdb := NewCdb(host, keyword)

	//构造插入语句
	// insert sql!
	insertsql := cdb.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		"me", gocql.TimeUUID(), "hello world")

	//执行,查看Exec方法,可知执行后已经关闭
	// exec a insert operation
	if err := insertsql.Exec(); err != nil {
		t.Fatal(err)
	}

	//构造查找语句
	// query sql
	querysql := cdb.Query(`SELECT "id", text FROM "tweet"`)

	//执行
	// done it !
	iter := querysql.Iter()

	//定义字节数组
	// id uuid in cassandra
	var id gocql.UUID
	var text string

	//循环取值，需要手工，无法再封装
	// take value just like that i can't wrap it again, and no need
	for iter.Scan(&id, &text) {
		t.Logf("Tweet:%v,%v\n", id, text)
	}

	//这个需要关闭
	// should be close
	if err := iter.Close(); err != nil {
		t.Logf("%v", err)
	}

}
