# i18n
golang i18n

## install

```shell
go get -u github.com/hiifong/i18n/v2
```

## Example

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hiifong/i18n/v2"
)

func main() {
	i18n.SetDefault(
		i18n.New(
			i18n.WithDir("/Users/hiifong/Desktop/Golang/i18n/example"),
			i18n.WithDefLang(i18n.ZhCN),
			i18n.WithLang(i18n.ZhCN),
			i18n.WithLang(i18n.EnUS),
		),
	)
	err := i18n.Load()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("all languages: %s\n", i18n.Languages())

	fmt.Println(i18n.Tr(i18n.EnUS, "hello"))
	fmt.Println(i18n.Tr(i18n.ZhCN, "hello_to", "hiifong"))
	fmt.Println(i18n.Tr(i18n.EnUS, "db.host"))
	fmt.Println(i18n.Tr(i18n.EnUS, "db.port"))
	fmt.Println(i18n.Tr(i18n.EnUS, "cache.redis.host"))
	fmt.Println(i18n.Tr(i18n.EnUS, "cache.redis.port"))
	fmt.Println(i18n.Tr(i18n.EnUS, "db.host"))
	fmt.Println(i18n.Tr(i18n.ZhCN, "db.port"))
	fmt.Println(i18n.Tr(i18n.ZhCN, "cache.redis.host"))
	fmt.Println(i18n.Tr(i18n.ZhCN, "cache.redis.port"))
	fmt.Println(i18n.Tr(i18n.ZhHK, "cache.redis.host"))

	fmt.Println("====================================")

	i18n.SetDefault(
		i18n.New(
			i18n.WithFS(http.Dir("/Users/hiifong/Desktop/Golang/i18n/example")),
			i18n.WithDefLang(i18n.EnUS),
			i18n.WithLang(i18n.ZhCN),
			i18n.WithLang(i18n.EnUS),
		),
	)
	err = i18n.Load()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("all languages: %s\n", i18n.Languages())
	fmt.Println(i18n.Tr(i18n.EnUS, "hello"))
	fmt.Println(i18n.Tr(i18n.ZhCN, "hello_to", "hiifong"))
	fmt.Println(i18n.Tr(i18n.EnUS, "db.host"))
	fmt.Println(i18n.Tr(i18n.EnUS, "db.port"))
	fmt.Println(i18n.Tr(i18n.EnUS, "cache.redis.host"))
	fmt.Println(i18n.Tr(i18n.EnUS, "cache.redis.port"))
	fmt.Println(i18n.Tr(i18n.EnUS, "db.host"))
	fmt.Println(i18n.Tr(i18n.ZhCN, "db.port"))
	fmt.Println(i18n.Tr(i18n.ZhCN, "cache.redis.host"))
	fmt.Println(i18n.Tr(i18n.ZhCN, "cache.redis.port"))
	fmt.Println(i18n.Tr(i18n.ZhHK, "cache.redis.host"))
}

// Output:
2025/05/13 22:24:24 loading file /Users/hiifong/Desktop/Golang/i18n/example/locale.zh-CN.ini
2025/05/13 22:24:24 loading file /Users/hiifong/Desktop/Golang/i18n/example/locale.en-US.ini
2025/05/13 22:24:24 loading file locale.zh-CN.ini
2025/05/13 22:24:24 loading file locale.en-US.ini
all languages: [zh-CN en-US]
hello
你好, hiifong
database host
database port
redis cache host
redis cache port
database host
数据库端口
redis缓存主机
redis缓存端口
redis缓存主机
====================================
all languages: [zh-CN en-US]
hello
你好, hiifong
database host
database port
redis cache host
redis cache port
database host
数据库端口
redis缓存主机
redis缓存端口
redis cache host
```
