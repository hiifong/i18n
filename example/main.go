package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hiifong/i18n/v2"
)

func main() {
	i18n.SetDefault(i18n.New(
		i18n.WithDir("./example"),
		i18n.WithAuto(),
	))
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

	i18n.SetDefault(i18n.New(
		i18n.WithFS(http.Dir("./example")),
		i18n.WithAuto(),
	))
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
