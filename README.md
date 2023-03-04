## 食用方法

这里的所有的文件内容，都是我靠调教chatgpt写出来的。

效果如

![](https://coding3min.oss-accelerate.aliyuncs.com/uPic/20230305/00-59-40-oe9tDL.png)

```go
├── README.md
├── files 文件上传位置
│   └── 111.png
├── go.mod
├── go.sum
├── main.go
└── templates 页面模板
    ├── download.html
    ├── error.html
    └── index.html
```

启动命令，使用的 80 端口。一定要用命令启动，不然找不到前端文件在哪。

```bash
go mod tidy
go run main.go
```

## 功能描述

* 使用 gin 做HTTP服务器。
* 使用os基本的上传下载处理。
* 基本的文件上传、下载、列表功能。

## 内网穿透

直接用此命令启动就行，用我的token。体验一下，有可能会出现自动切换token的情况，建议自己申请。

```bash
chmod 755 ./natapp
./natapp -authtoken e8f05cbd1e1b139d
```

windows 双击允许就行了。自己申请修改一下`config.ini`

自己申请在 https://natapp.cn/
