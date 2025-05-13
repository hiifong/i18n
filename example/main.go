package main

import (
	"fmt"
	"log"

	"github.com/hiifong/i18n"
)

func main() {
	i18n.SetDefault(
		i18n.New(
			i18n.WithDir("D:\\hiifong\\Desktop\\project\\i18n\\example"),
			i18n.WithLang("zh-CN"),
		),
	)
	err := i18n.Load()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(i18n.Tr("en-US", "hello"))
	fmt.Println(i18n.Tr("zh-CN", "hello_to", "hiifong"))
	fmt.Println(i18n.Tr("en-US", "db.host"))
	fmt.Println(i18n.Tr("en-US", "db.port"))
	fmt.Println(i18n.Tr("en-US", "cache.redis.host"))
	fmt.Println(i18n.Tr("en-US", "cache.redis.port"))
	fmt.Println(i18n.Tr("en-US", "db.host"))
	fmt.Println(i18n.Tr("zh-CN", "db.port"))
	fmt.Println(i18n.Tr("zh-CN", "cache.redis.host"))
	fmt.Println(i18n.Tr("zh-CN", "cache.redis.port"))
	fmt.Println(i18n.Tr("zh-HK", "cache.redis.host"))
}
