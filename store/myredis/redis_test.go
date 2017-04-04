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
package myredis

import (
	"testing"
)

func TestRedis(t *testing.T) {
	config := RedisConfig{}
	config.DB = 0
	config.Host = "127.0.0.1:6379"
	config.Password = "smart2016" // no password set ""

	client, err := NewRedis(config) // new redis client
	if err != nil {
		panic(err)
	}

	// set key==value
	err = client.Set("key", "value", 0)
	if err != nil {
		t.Error(err.Error())
	}

	// get key
	val, err := client.Get("key")
	if err != nil {
		panic(err)
	} else {
		t.Log("Redis value:" + val)
	}

	// push test,pust pool with b value
	num, err := client.Lpush("pool", "b")
	if err != nil {
		t.Error(err.Error())
	}

	// total length of list
	t.Log(num)

	// pushx test,will be error if not exist pool10
	num, err = client.Lpushx("pool10", "b")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(num)

	// len test
	t.Log(client.Llen("pool"))

	// pop test
	pops, e := client.Lpop("pool2")
	if e != nil {
		t.Logf("%v,%v", pops, e)
	}

	// bpop test
	bpops, e := client.Blpop(2, "pool1", "pool1")
	t.Logf("%#v,%v", bpops, e)

	//rpoplpush test POOL1 empty so will be redis.nil
	rpoplpush, e := client.Rpoplpush("POOL1", "pool1")
	t.Logf("%#v,%v", rpoplpush, e)

	//brpoplpush test POOL1 empty so will be redis.nil if timeout but zero set is wait a long time
	brpoplpush, e := client.Brpoplpush("POOL1", "pool1", 2)
	t.Logf("%#v,%v", brpoplpush, e)
}
