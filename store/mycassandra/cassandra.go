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

/*
cassandra操作工具类
*/
import (
	"github.com/gocql/gocql"
)

//配置特别简单，主机名和keyspace
// cassand config,just need host and keyspace
type Config struct {
	Host     []string
	Keyspace string
}

//这是一个重要的结构体，cassandra连接包装于此
// cassandra client
type Cdb struct {
	*Config
	baseSession *gocql.Session
}

//配置初始化连接
// init cassandra by config
func NewCdbWithConf(c *Config) (cdb *Cdb) {
	cdb = &Cdb{
		Config: c,
	}
	cdb.Connect()
	return
}

//一般初始化连接
// init cdb just these
func NewCdb(host []string, keyspace string) (cdb *Cdb) {
	cdb = &Cdb{
		Config: &Config{
			Host:     host,
			Keyspace: keyspace,
		},
	}
	cdb.Connect()
	return
}

//连接，出错panic
// before use must connect,error will be panic,oh!
func (self *Cdb) Connect() {
	cluster := gocql.NewCluster(self.Host...)
	cluster.Keyspace = self.Keyspace
	cluster.Consistency = gocql.Quorum
	baseSession, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	self.baseSession = baseSession

}

//别名，为了容易使用
//alias name easy use
func New(c *Config) (cdb *Cdb) {
	cdb = &Cdb{
		Config: c,
	}
	cdb.Connect()
	return
}

//构造查询语句，包括插入，查找，不仅仅是查询
//create a cassnara sql,not just for query can be insert
func (c *Cdb) Query(stmt string, values ...interface{}) *gocql.Query {
	if c.baseSession == nil {
		return nil
	}
	return c.baseSession.Query(stmt, values...)
}

//使用上面的查询语句，开始执行,一般是插入操作
// Query() then can Exec just for insert
func (c *Cdb) Exec(q *gocql.Query) error {
	return q.Exec()
}

//使用上面的查询语句，开始执行,一般是查找操作
// Query() then real query,see test
func (c *Cdb) Iter(q *gocql.Query) *gocql.Iter {
	return q.Iter()
}
