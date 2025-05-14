package i18n

import (
	"fmt"
	"html/template"
)

type Interface interface {
	Load() error
	Default() Lang
	SetDefault(Lang)
	Languages() []Lang
	Tr(lang Lang, key string, values ...any) template.HTML
}

var (
	defI18n Interface
)

func SetDefault(i Interface) {
	if i != nil {
		defI18n = i
	}
}

func Load() error {
	if defI18n == nil {
		return fmt.Errorf("no default i18n")
	}
	return defI18n.Load()
}

func Default() Lang {
	return defI18n.Default()
}

func SetDefaultLang(lang Lang) {
	defI18n.SetDefault(lang)
}

func Tr(lang Lang, key string, values ...any) template.HTML {
	if defI18n == nil {
		return ""
	}
	return defI18n.Tr(lang, key, values...)
}

func Languages() []Lang {
	if defI18n == nil {
		return nil
	}
	return defI18n.Languages()
}
