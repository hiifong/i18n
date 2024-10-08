package i18n

import (
	"fmt"
	"sync"
)

type defaultI18n struct {
	mu    sync.RWMutex
	first string
	i18n  map[string]map[string]Language
}

type defaultOption func(*defaultI18n)

func defaultAdapterWithDefaultLang(lang string) defaultOption {
	return func(i *defaultI18n) {
		i.first = lang
	}
}

func defaultAdapterWithLang(lang string, i18n interface{}) defaultOption {
	return func(i *defaultI18n) {
		i.i18n[lang] = i18n.(map[string]Language)
	}
}

func newDefault(options ...defaultOption) (*defaultI18n, error) {
	i := new(defaultI18n)
	i.i18n = make(map[string]map[string]Language)
	for _, option := range options {
		if option != nil {
			option(i)
		}
	}
	return i, nil
}

var _ adapter = (*defaultI18n)(nil)

func (i *defaultI18n) register(lang string, i18n interface{}) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	if lang == "" {
		return fmt.Errorf("lang can't be empty")
	}
	if _, ok := i.i18n[lang]; ok {
		return fmt.Errorf("language %s is already registered", lang)
	}
	l, ok := i18n.(map[string]Language)
	if !ok {
		return fmt.Errorf("this %+v is not support", lang)
	}
	i.i18n[lang] = l
	return nil
}

func (i *defaultI18n) update(lang, key string, i18n interface{}) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	if lang == "" {
		return fmt.Errorf("language can't be empty")
	}
	if _, ok := i.i18n[lang]; !ok {
		return fmt.Errorf("language %s is not registered", lang)
	}
	l, ok := i18n.(Language)
	if !ok {
		return fmt.Errorf("this %s is not support", lang)
	}
	if key != l.Key {
		return fmt.Errorf("this %s key is not match", lang)
	}
	i.i18n[lang][key] = l
	return nil
}

func (i *defaultI18n) setDefault(lang string) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	if lang == "" {
		return fmt.Errorf("lang can't be empty")
	}
	if _, ok := i.i18n[lang]; !ok {
		return fmt.Errorf("language %s is not registered", lang)
	}
	i.first = lang
	return nil
}

func (i *defaultI18n) t(lang string, key string) (string, string, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	if i.first == "" {
		return key, "", fmt.Errorf("default language is not set")
	}
	t := func(lang string, key string) (string, string, error) {
		lm, ok := i.i18n[lang]
		if !ok {
			lm, ok = i.i18n[i.first]
		}
		l, ok := lm[key]
		if !ok {
			l, ok = i.i18n[i.first][key]
		}
		if !ok {
			return key, "", fmt.Errorf("language %s is not registered", lang)
		}
		if l.Raw == "" {
			return key, "", fmt.Errorf("language %s is empty", lang)
		}
		if key != l.Key {
			return key, "", fmt.Errorf("language %s key is not match", lang)
		}
		return l.Key, l.Raw, nil
	}

	if lang == "" {
		return t(i.first, key)
	}
	return t(lang, key)
}

func (i *defaultI18n) onlyT(lang string, key string) string {
	_, text, err := i.t(lang, key)
	if err != nil {
		return ""
	}
	return text
}
