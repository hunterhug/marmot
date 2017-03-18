# 下载任意网站图片

写入taobao.csv：

```
https://detail.tmall.com/item.htm?id=523350171126&skuId=3120562159704,tmall
#https://item.taobao.com/item.htm?id=40066362090,taobao
http://pic.yesky.com/492/97556992.shtml,meitu
```

链接分为两部分，前面是链接，后面是图片保存的目录名，`#`表示忽略这一个网站

跑起来，-config后面是taobao.csv的位置,如果在/data/app下，那么需-config=/data/app/taobao.csv

```
go run taobao.go -config=taobao.csv
taobao.exe -config=taobao.csv
```

![](see.png)
![](pic.png)