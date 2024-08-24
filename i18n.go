package i18n

import (
	"fmt"
)

// I18ner 国际化接口
type I18ner interface {
	// Register 注册新的语言
	Register(lang string, i18n interface{}) error

	// Add 添加翻译
	Add(lang string, code int, t interface{}) error

	// SetDefault 设置默认语言
	SetDefault(lang string) error

	// T 获取翻译
	T(lang string, code int) (int, string, error)

	// OnlyT 仅获取翻译
	OnlyT(lang string, code int) string
}

type Language struct {
	Code int
	Raw  interface{}
}

func (l *Language) String() string {
	return fmt.Sprintf("Language {Code: %d, Raw: %v}", l.Code, l.Raw)
}

type I18n struct {
	first string
	i18n  map[string]map[int]Language
}

type Option func(*I18n)

func WithDefaultLang(lang string) Option {
	return func(i *I18n) {
		i.first = lang
	}
}

func WithLang(lang string, i18n interface{}) Option {
	return func(i *I18n) {
		i.i18n[lang] = i18n.(map[int]Language)
	}
}

func New(options ...Option) *I18n {
	i18n := new(I18n)
	i18n.i18n = make(map[string]map[int]Language)
	for _, option := range options {
		if option != nil {
			option(i18n)
		}
	}
	return i18n
}

var _ I18ner = (*I18n)(nil)

func (i *I18n) Register(lang string, i18n interface{}) error {
	if _, ok := i.i18n[lang]; ok {
		return fmt.Errorf("language %s is already registered", lang)
	}
	l, ok := i18n.(map[int]Language)
	if !ok {
		return fmt.Errorf("this %+v is not support", lang)
	}
	i.i18n[lang] = l
	return nil
}

func (i *I18n) Add(lang string, code int, t interface{}) error {
	if _, ok := i.i18n[lang]; !ok {
		return fmt.Errorf("language %s is not registered", lang)
	}
	l, ok := t.(Language)
	if !ok {
		return fmt.Errorf("this %s is not support", lang)
	}
	if code != l.Code {
		return fmt.Errorf("this %s code is not match", lang)
	}
	i.i18n[lang][code] = l
	return nil
}

func (i *I18n) SetDefault(lang string) error {
	if _, ok := i.i18n[lang]; !ok {
		return fmt.Errorf("language %s is not registered", lang)
	}
	i.first = lang
	return nil
}

func (i *I18n) T(lang string, code int) (int, string, error) {
	if i.first == "" {
		return code, "", fmt.Errorf("default language is not set")
	}
	t := func(lang string, code int) (int, string, error) {
		l, ok := i.i18n[lang][code]
		if !ok {
			return code, "", fmt.Errorf("language %s is not registered", lang)
		}
		msg, ok := l.Raw.(string)
		if !ok {
			return code, "", fmt.Errorf("language %s is not support", lang)
		}
		if msg == "" {
			return code, "", fmt.Errorf("language %s is empty", lang)
		}
		return code, msg, nil
	}

	if lang == "" {
		return t(i.first, code)
	}
	return t(lang, code)
}

func (i *I18n) OnlyT(lang string, code int) string {
	_, msg, err := i.T(lang, code)
	if err != nil {
		return ""
	}
	return msg
}
