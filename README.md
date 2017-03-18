# Golang Spider

## 一.下载

自己封装的爬虫库，类似于Python的requests，又不像，你只需通过该方式获取库

```
go get -u -v github.com/hunterhug/GoSpider
```

或者新建 你的GOPATH路径/src/github.com/hunterhug

```
cd src/github.com/hunterhug
git clone https://github.com/hunterhug/GoSpider
```

此库采用[Glide](https://github.com/Masterminds/glide)方式管理第三方库（贡献者可以查看）

```
$ glide init                              # 创建工作区
$ open glide.yaml                         # 编辑glide.yaml文件
$ glide get github.com/Masterminds/cookoo # get下库然后会自动写入glide.yaml
$ glide install                           # 安装,没有glide.lock,会先运行glide up

# work, work, work
$ go build                                # 试试可不可以跑
$ glide up                                # 更新库，创建glide.lock
```

默认所有第三方库已经保存在vendor

## 二.使用

## 例子
1.任意图片下载,见[例子1](example/taobao/README.md)

## 三.记录
20170318 

1.新增glide管理第三方库
2.更新若干函数
3.修改README.md等
4.增加任意图片下载示例


## 联系

QQ：569929309 一只尼玛
公众号：lenggirlcom(搬砖的陈大师)

# LICENSE

欢迎加功能

```
Copyright 2017 hunterhug/一只尼玛.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License
```