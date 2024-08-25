package i18n

import (
	"github.com/hiifong/i18n/ent/schema"
)

type EntI18n struct {
	schema.EntI18n
}

type entOption func(*EntI18n)

func entAdapterWithDefaultLang(lang string) entOption {
	return func(i *EntI18n) {
		// TODO
	}
}

func entAdapterWithLang(lang string, i18n interface{}) entOption {
	return func(i *EntI18n) {
		// TODO
	}
}

func newEnt(dns string, options ...entOption) (*EntI18n, error) {
	i := new(EntI18n)
	for _, option := range options {
		if option != nil {
			option(i)
		}
	}
	return i, nil
}

var _ adapter = (*EntI18n)(nil)

func (i *EntI18n) register(lang string, i18n interface{}) error {
	// TODO
	return nil
}

func (i *EntI18n) update(lang, key string, i18n interface{}) error {
	// TODO
	return nil
}

func (i *EntI18n) setDefault(lang string) error {
	// TODO
	return nil
}

func (i *EntI18n) t(lang string, key string) (string, string, error) {
	// TODO
	return "", "", nil
}

func (i *EntI18n) onlyT(lang string, key string) string {
	// TODO
	return ""
}
