package i18n

import "github.com/hiifong/i18n/ent/schema"

type entI18n struct {
	schema.EntI18n
}

type entOption func(*entI18n)

func entAdapterWithDefaultLang(lang string) entOption {
	return func(i *entI18n) {
		// TODO
	}
}

func entAdapterWithLang(lang string, i18n interface{}) entOption {
	return func(i *entI18n) {
		// TODO
	}
}

func newEnt(dns string, options ...entOption) (*entI18n, error) {
	i := new(entI18n)
	for _, option := range options {
		if option != nil {
			option(i)
		}
	}
	return i, nil
}

var _ adapter = (*entI18n)(nil)

func (i *entI18n) register(lang string, i18n interface{}) error {
	// TODO
	return nil
}

func (i *entI18n) update(lang, key string, i18n interface{}) error {
	// TODO
	return nil
}

func (i *entI18n) setDefault(lang string) error {
	// TODO
	return nil
}

func (i *entI18n) t(lang string, key string) (string, string, error) {
	// TODO
	return "", "", nil
}

func (i *entI18n) onlyT(lang string, key string) string {
	// TODO
	return ""
}
