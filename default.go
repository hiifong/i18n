package i18n

import (
	"fmt"
	"sync"
)

type DefaultAdapter struct {
	mu    sync.RWMutex
	first string
	i18n  map[string]map[string]Language
}

type DefaultOption func(*DefaultAdapter)

func DefaultAdapterWithDefaultLang(lang string) DefaultOption {
	return func(i *DefaultAdapter) {
		i.first = lang
	}
}

func DefaultAdapterWithLang(lang string, i18n interface{}) DefaultOption {
	return func(i *DefaultAdapter) {
		i.i18n[lang] = i18n.(map[string]Language)
	}
}

func NewDefault(options ...DefaultOption) (*DefaultAdapter, error) {
	i := new(DefaultAdapter)
	i.i18n = make(map[string]map[string]Language)
	for _, option := range options {
		if option != nil {
			option(i)
		}
	}
	return i, nil
}

func (i *DefaultAdapter) Register(lang string, i18n interface{}) error {
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

func (i *DefaultAdapter) Update(lang, key string, i18n interface{}) error {
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

func (i *DefaultAdapter) SetDefault(lang string) error {
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

func (i *DefaultAdapter) T(lang string, key string) (string, string, error) {
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

func (i *DefaultAdapter) OnlyT(lang string, key string) string {
	_, msg, err := i.T(lang, key)
	if err != nil {
		return ""
	}
	return msg
}
