package main

import (
	"fmt"

	"github.com/hiifong/i18n"
)

func main() {
	zhCN := map[string]i18n.Language{
		"hello": {
			Key: "hello",
			Raw: "你好",
		},
		"world": {
			Key: "world",
			Raw: "世界",
		},
		"hello_world": {
			Key: "hello_world",
			Raw: "你好, 世界",
		},
	}
	enUS := map[string]i18n.Language{
		"hello": {
			Key: "hello",
			Raw: "Hello",
		},
		"world": {
			Key: "world",
			Raw: "World",
		},
		"hello_world": {
			Key: "hello_world",
			Raw: "Hello, World",
		},
	}
	i, _ := i18n.New(
		i18n.WithAdapter(i18n.Default),
		i18n.WithDefaultLang(i18n.EnUS),
		i18n.WithLang(i18n.ZhCN, zhCN),
	)
	err := i.Register(i18n.ZhCN, zhCN)
	if err != nil {
		fmt.Println(err)
	}
	err = i.Register(i18n.EnUS, enUS)
	if err != nil {
		fmt.Println(err)
	}
	code, msg, err := i.T(i18n.ZhCN, "hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("key: %s, msg: %s\n", code, msg)

	code, msg, err = i.T("", "hello_world")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("key: %s, msg: %s\n", code, msg)

	err = i.Update(i18n.EnUS, "world", i18n.Language{
		Key: "world",
		Raw: "Hello, World",
	})
	if err != nil {
		fmt.Println(err)
	}
	key, msg, err := i.T(i18n.EnUS, "world")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("key: %s, msg: %s\n", key, msg)

	err = i.Update(i18n.ZhCN, "hello", nil)
	if err != nil {
		fmt.Println(err)
	}

	err = i.Update(i18n.ZhCN, "unknown", i18n.Language{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("onlye t: %s\n", i.OnlyT(i18n.ZhCN, "hello_world"))

	key, msg, err = i.T(i18n.ZhHK, "hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("zh_HK key: %s --> msg: %s --> err:%v\n", key, msg, err)
}
