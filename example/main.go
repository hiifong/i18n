package main

import (
	"fmt"
	"log"
	"os"

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
	fmt.Println("en-US:")
	fmt.Println(i18n.Tr(i18n.EnUS, "hello_to", "hiifong"))
	fmt.Println(i18n.Tr(i18n.EnUS, "i_have_1_apple"))
	fmt.Println(i18n.Tr(i18n.EnUS, "i_have_n_apples", 100))
	fmt.Println(i18n.Tr(i18n.EnUS, "about.name", "hiifong"))
	fmt.Println(i18n.Tr(i18n.EnUS, "about.email", "i@hiif.ong"))
	fmt.Println(i18n.Tr(i18n.EnUS, "about.github", "https://github.com/hiifong"))
	fmt.Println(i18n.Tr(i18n.EnUS, "about.website", "https://github.com/hiifong"))
	fmt.Println("-----------------------------------------------------")

	i18n.SetDefault(i18n.New(
		i18n.WithFS(os.DirFS("./example")),
		i18n.WithAuto(),
	))
	err = i18n.Load()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("all languages: %s\n", i18n.Languages())
	fmt.Println("zh-CN:")
	fmt.Println(i18n.Tr(i18n.ZhCN, "hello_to", "hiifong"))
	fmt.Println(i18n.Tr(i18n.ZhCN, "i_have_1_apple"))
	fmt.Println(i18n.Tr(i18n.ZhCN, "i_have_n_apples", 100))
	fmt.Println(i18n.Tr(i18n.ZhCN, "about.name", "hiifong"))
	fmt.Println(i18n.Tr(i18n.ZhCN, "about.email", "i@hiif.ong"))
	fmt.Println(i18n.Tr(i18n.ZhCN, "about.github", "https://github.com/hiifong"))
	fmt.Println(i18n.Tr(i18n.ZhCN, "about.website", "https://github.com/hiifong"))
}
