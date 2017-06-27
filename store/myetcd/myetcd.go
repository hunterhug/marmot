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
package myetcd

import (
	"fmt"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	//"log"
	"errors"
	"strings"
	"time"
)

type EtcdConfig struct {
	Host    string
	Port    string
	Prefix  string
	Timeout int
}

type MyEtcd struct {
	Config EtcdConfig
	Client client.KeysAPI
}

// 使用见test
func NewEtcd(config EtcdConfig) (*MyEtcd, error) {
	myetcd := &MyEtcd{Config: config}
	var etcdUrls = []string{}
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Port == "" {
		config.Port = "2379"
	}
	for _, url := range strings.Split(config.Host, ",") {
		etcdUrl := strings.TrimSpace(url)
		etcdUrl = fmt.Sprintf("http://%s:%s", etcdUrl, config.Port)
		etcdUrls = append(etcdUrls, etcdUrl)
	}

	cfg := client.Config{
		Endpoints: etcdUrls,
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	myetcd.Client = client.NewKeysAPI(c)
	return myetcd, nil
}

//  放值，是目录会失败
func (myetcd *MyEtcd) Set(key, value string) error {
	key = fmt.Sprintf("%s/%s", myetcd.Config.Prefix, key)
	_, err := myetcd.Client.Set(context.Background(), key, value, nil)
	//log.Printf("%#v", resp)
	return err
}

//  创建目录，存在Key报错，目录已经存在不报错,如果a/b是key，那么不能新建a/b/c目录
func (myetcd *MyEtcd) SetDir(key string) error {
	key = fmt.Sprintf("%s/%s", myetcd.Config.Prefix, key)
	resp, err := myetcd.Client.Get(context.Background(), key, nil)
	// 存在
	if err == nil {
		if resp.Node.Dir {
			return nil
		}
		return errors.New(fmt.Sprintf("Is a key not dir(%s)", key))
	}
	_, err = myetcd.Client.Set(context.Background(), key, "dir", &client.SetOptions{Dir: true})
	//log.Printf("%#v", resp)
	return err
}

// 严格模式取，不能取目录
func (myetcd *MyEtcd) StrictGet(key string) (string, error) {
	key = fmt.Sprintf("%s/%s", myetcd.Config.Prefix, key)
	resp, err := myetcd.Client.Get(context.Background(), key, nil)
	if err != nil {
		return "", err
	} else {
		if resp.Node.Dir {
			return "", errors.New(fmt.Sprintf("Get a dir (%s)", key))
		}
		// print common key info
		//log.Printf("Get is done. Metadata is %q\n", resp)
		// print value
		//log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}
	return resp.Node.Value, err
}

// 可以取目录：空
func (myetcd *MyEtcd) Get(key string) string {
	key = fmt.Sprintf("%s/%s", myetcd.Config.Prefix, key)
	resp, err := myetcd.Client.Get(context.Background(), key, nil)
	if err != nil {
		//log.Println(err.Error())
		return ""
	}
	return resp.Node.Value
}

// 是否存在Key
func (myetcd *MyEtcd) Exist(key string) bool {
	key = fmt.Sprintf("%s/%s", myetcd.Config.Prefix, key)
	_, err := myetcd.Client.Get(context.Background(), key, nil)
	if err == nil {
		return true
	}
	return false
}

// 是否是目录，nil则是
func (myetcd *MyEtcd) IsDir(key string) error {
	key = fmt.Sprintf("%s/%s", myetcd.Config.Prefix, key)
	resp, err := myetcd.Client.Get(context.Background(), key, nil)
	if err != nil {
		return err
	}
	if !resp.Node.Dir {
		return errors.New("Is a key not dir")
	}
	return nil
}

// 删除key
func (myetcd *MyEtcd) Rm(key string) error {
	key = fmt.Sprintf("%s/%s", myetcd.Config.Prefix, key)
	_, err := myetcd.Client.Delete(context.Background(), key, nil)
	if err != nil {
		if client.IsKeyNotFound(err) {
			return nil
		}
	}
	return err
}

func (myetcd *MyEtcd) StrictRm(key string) error {
	key = fmt.Sprintf("%s/%s", myetcd.Config.Prefix, key)
	_, err := myetcd.Client.Delete(context.Background(), key, nil)
	return err
}

func (myetcd *MyEtcd) IsKeyNotFound(e error) bool {
	return client.IsKeyNotFound(e)
}

// 级联删除key（值也是）
func (myetcd *MyEtcd) RmAll(key string) error {
	key = fmt.Sprintf("%s/%s", myetcd.Config.Prefix, key)
	_, err := myetcd.Client.Delete(context.Background(), key, &client.DeleteOptions{Recursive: true})
	if err != nil {
		if client.IsKeyNotFound(err) {
			return nil
		}
	}
	return err
}

// 列出儿子们
func (myetcd *MyEtcd) List(key string) (client.Nodes, error) {
	key = fmt.Sprintf("%s/%s", myetcd.Config.Prefix, key)
	resp, err := myetcd.Client.Get(context.Background(), key, nil)
	if err != nil {
		return nil, err
	}
	if resp.Node.Dir {
		return resp.Node.Nodes, nil
	}
	return nil, errors.New(fmt.Sprintf("Is is a key not dir(%s)", key))
}

func (myetcd *MyEtcd) ListR(key string) (client.Nodes, error) {
	key = fmt.Sprintf("%s/%s", myetcd.Config.Prefix, key)
	resp, err := myetcd.Client.Get(context.Background(), key, &client.GetOptions{Recursive: true})
	if err != nil {
		return nil, err
	}
	if resp.Node.Dir {
		return resp.Node.Nodes, nil
	}
	return nil, errors.New(fmt.Sprintf("Is is a key not dir(%s)", key))
}

func (myetcd *MyEtcd) GetAll(key string) (*client.Node, error) {
	key = fmt.Sprintf("%s/%s", myetcd.Config.Prefix, key)
	resp, err := myetcd.Client.Get(context.Background(), key, &client.GetOptions{Recursive: true})
	if err != nil {
		return nil, err
	}
	return resp.Node, nil
}
