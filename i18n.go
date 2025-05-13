package i18n

import (
	"fmt"
)

type Interface interface {
	Load() error
	Tr(lang Lang, key string, values ...any) string
	Languages() []Lang
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

func Tr(lang Lang, key string, values ...any) string {
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
