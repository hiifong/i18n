package i18n

import (
	"github.com/hiifong/i18n/ent/schema"
)

type EntI18n struct {
	schema.EntI18n
}

type EntOption func(n *EntI18n)

func EntAdapterWithDefaultLang(lang string) EntOption {
	return func(i *EntI18n) {
		// TODO
	}
}

func EntAdapterWithLang(lang string, i18n interface{}) EntOption {
	return func(i *EntI18n) {
		// TODO
	}
}

func NewEnt(dns string, options ...EntOption) (*EntI18n, error) {
	i := new(EntI18n)
	for _, option := range options {
		if option != nil {
			option(i)
		}
	}
	return i, nil
}

var _ I18ner = (*EntI18n)(nil)

func (i *EntI18n) Register(lang string, i18n interface{}) error {
	// TODO
	return nil
}

func (i *EntI18n) Update(lang, key string, i18n interface{}) error {
	// TODO
	return nil
}

func (i *EntI18n) SetDefault(lang string) error {
	// TODO
	return nil
}

func (i *EntI18n) T(lang string, key string) (string, string, error) {
	// TODO
	return "", "", nil
}

func (i *EntI18n) OnlyT(lang string, key string) string {
	// TODO
	return ""
}
