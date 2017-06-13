package myetcd

import (
	"fmt"
	"testing"
)

func TestNewEtcd(t *testing.T) {
	etcdConfig := EtcdConfig{
		Host:    "127.0.0.1",
		Prefix:  "a/b",
		Timeout: 1,
	}
	ec, err := NewEtcd(etcdConfig)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		err = ec.RmAll("")
		if err != nil {
			fmt.Println("rmall err:" + err.Error())
		}
		key := "c"
		dir := "c/d"

		// 两次
		err = ec.Set(key, "ff")
		if err != nil {
			fmt.Println("1 set err:" + err.Error())
		}
		err = ec.Set(key, "ff")
		if err != nil {
			fmt.Println("2 set err:" + err.Error())
		}

		// 两次 失败，
		err = ec.SetDir(dir)
		if err != nil {
			fmt.Println("1 set dir err:" + err.Error())
		}
		err = ec.SetDir(dir)
		if err != nil {
			fmt.Println("2 set dir err:" + err.Error())
		}

		err = ec.Set(dir, "ff")
		if err != nil {
			fmt.Println("3 set err:" + err.Error())
		}

		// key已经存在
		err = ec.SetDir(key)
		if err != nil {
			fmt.Println("3 set dir err:" + err.Error())
		}

		r := ec.Get(key)
		fmt.Println(r)

		r, err = ec.StrictGet(dir)
		fmt.Printf("get %v,error:%v:exist?%v\n", r, err, ec.Exist(dir))

		err = ec.IsDir(dir)
		fmt.Printf("1 isdir:%v\n", err)

		err = ec.IsDir(key)
		fmt.Printf("2 isdir:%v\n", err)

		err = ec.Set(dir+"/dd", "ff")
		if err != nil {
			fmt.Println("4 set err:" + err.Error())
		}
		err = ec.Rm(dir)
		fmt.Printf("%v\n", err)
	}
}

func TestNewEtcd2(t *testing.T) {

	etcdConfig := EtcdConfig{
		Host:    "127.0.0.1",
		Prefix:  "a/b",
		Timeout: 1,
	}
	ec, err := NewEtcd(etcdConfig)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		err = ec.RmAll("")
		if err != nil {
			fmt.Println("rmall err:" + err.Error())
		}
		key := "c"
		dir := "d"

		// 两次
		err = ec.Set(key, "ff")
		if err != nil {
			fmt.Println("1 set err:" + err.Error())
		}
		err = ec.Set(key, "ff")
		if err != nil {
			fmt.Println("2 set err:" + err.Error())
		}

		// 两次 失败，
		err = ec.SetDir(dir)
		if err != nil {
			fmt.Println("1 set dir err:" + err.Error())
		}
		err = ec.SetDir(dir)
		if err != nil {
			fmt.Println("2 set dir err:" + err.Error())
		}

		err = ec.Set(dir, "ff")
		if err != nil {
			fmt.Println("3 set err:" + err.Error())
		}

		// key已经存在
		err = ec.SetDir(key)
		if err != nil {
			fmt.Println("3 set dir err:" + err.Error())
		}

		r := ec.Get(key)
		fmt.Println(r)

		r, err = ec.StrictGet(dir)
		fmt.Printf("get %v,error:%v:exist?%v\n", r, err, ec.Exist(dir))

		err = ec.IsDir(dir)
		fmt.Printf("1 isdir:%v\n", err)

		err = ec.IsDir(key)
		fmt.Printf("2 isdir:%v\n", err)

		err = ec.Set(dir+"/dd", "ff")
		if err != nil {
			fmt.Println("4 set err:" + err.Error())
		}
		err = ec.Rm(dir)
		fmt.Printf("rm dir %v\n", err)

		err = ec.RmAll(key)
		fmt.Printf("rm dir way rm key:%v\n", err)
		err = ec.RmAll("")
		if err != nil {
			fmt.Println("rmall err:" + err.Error())
		}

		ec.SetDir("ff")
		ec.SetDir("ff/dd")
		ec.SetDir("ff/dd/dd")
		ec.Set("ff/df", "ddd")
		node, e := ec.List("ff")
		if e != nil {
			fmt.Println(e.Error())
		} else {
			fmt.Printf("%#v:%v\n", node, len(node))
			for _, n := range node {
				fmt.Println(n.Key, n.Value, n.Dir)
			}
		}
	}
}
