package i18n

import (
	"fmt"
	"sync"
)

// I18ner 国际化接口
type I18ner interface {
	// Register 注册新的语言
	Register(lang string, i18n interface{}) error

	// Update 更新翻译, 如果存在翻译则更新，否则添加翻译
	Update(lang, key string, i18n interface{}) error

	// SetDefault 设置默认语言
	SetDefault(lang string) error

	// T 获取翻译
	T(lang, key string) (string, string, error)

	// OnlyT 仅获取翻译
	OnlyT(lang string, key string) string
}

type Language struct {
	Key, Raw string
}

func (l *Language) String() string {
	return fmt.Sprintf("Language {Key: %s, Raw: %v}", l.Key, l.Raw)
}

type I18n struct {
	mu    sync.RWMutex
	first string
	i18n  map[string]map[string]Language
}

type Option func(*I18n)

func WithDefaultLang(lang string) Option {
	return func(i *I18n) {
		i.first = lang
	}
}

func WithLang(lang string, i18n interface{}) Option {
	return func(i *I18n) {
		i.i18n[lang] = i18n.(map[string]Language)
	}
}

func New(options ...Option) *I18n {
	i18n := new(I18n)
	i18n.i18n = make(map[string]map[string]Language)
	for _, option := range options {
		if option != nil {
			option(i18n)
		}
	}
	return i18n
}

var _ I18ner = (*I18n)(nil)

func (i *I18n) Register(lang string, i18n interface{}) error {
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

func (i *I18n) Update(lang, key string, i18n interface{}) error {
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

func (i *I18n) SetDefault(lang string) error {
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

func (i *I18n) T(lang string, key string) (string, string, error) {
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

func (i *I18n) OnlyT(lang string, key string) string {
	_, msg, err := i.T(lang, key)
	if err != nil {
		return ""
	}
	return msg
}
