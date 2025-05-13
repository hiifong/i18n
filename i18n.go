package i18n

import "sync"

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
	return defI18n.Load()
}

func Tr(lang Lang, key string, values ...any) string {
	return defI18n.Tr(lang, key, values...)
}

func Languages() []Lang {
	return defI18n.Languages()
}
