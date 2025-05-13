package i18n

import (
	"fmt"
	"sync"
)

type Interface interface {
	Load() error
	Tr(lang Lang, key string, values ...any) string
	Languages() []Lang
}

var (
	defI18nOnce sync.Once
	defI18n     Interface
)

func Default() Interface {
	defI18nOnce.Do(func() {
		if defI18n == nil {
			defI18n = nil
		}
	})
	return defI18n
}

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
