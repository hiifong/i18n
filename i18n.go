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

type Type string

const (
	Default Type = "default"
	Ent     Type = "ent"
	Gorm    Type = "gorm"
)

var (
	once = sync.Once{}
)

type Language struct {
	Key, Raw string
}

func (l *Language) String() string {
	return fmt.Sprintf("Language {Key: %s, Raw: %v}", l.Key, l.Raw)
}

type I18n struct {
	t       Type // adapter type
	dns     string
	opts    []interface{}
	adapter I18ner
}

type Option func(*I18n)

func WithAdapter(t Type, dns ...string) Option {
	return func(i *I18n) {
		if t != "" {
			i.t = t
		}
		if len(dns) == 1 {
			i.dns = dns[0]
		}
	}
}

func WithDefaultLang(lang string) Option {
	return func(i *I18n) {
		switch i.t {
		case Ent:
			i.opts = append(i.opts, EntAdapterWithDefaultLang(lang))
		case Gorm:
			i.opts = append(i.opts, GormAdapterWithDefaultLang(lang))
		case Default:
			i.opts = append(i.opts, DefaultAdapterWithDefaultLang(lang))
		default:
			i.opts = append(i.opts, DefaultAdapterWithDefaultLang(lang))
		}
	}
}

func WithLang(lang string, i18n interface{}) Option {
	return func(i *I18n) {
		switch i.t {
		case Ent:
			i.opts = append(i.opts, EntAdapterWithLang(lang, i18n))
		case Gorm:
			i.opts = append(i.opts, GormAdapterWithLang(lang, i18n))
		case Default:
			i.opts = append(i.opts, DefaultAdapterWithLang(lang, i18n))
		default:
			i.opts = append(i.opts, DefaultAdapterWithLang(lang, i18n))
		}
	}
}

func New(options ...Option) (*I18n, error) {
	i := new(I18n)
	var err error
	//once.Do(func() {
	i.t = Default
	for _, option := range options {
		if option != nil {
			option(i)
		}
	}
	switch i.t {
	case Ent:
		var opts []EntOption
		for _, opt := range i.opts {
			opts = append(opts, opt.(EntOption))
		}
		i.adapter, err = NewEnt(i.dns, opts...)
	case Gorm:
		var opts []GormOption
		for _, opt := range i.opts {
			opts = append(opts, opt.(GormOption))
		}
		i.adapter, err = NewGorm(i.dns, opts...)
	default:
		var opts []DefaultOption
		for _, opt := range i.opts {
			opts = append(opts, opt.(DefaultOption))
		}
		i.adapter, err = NewDefault(opts...)
	}
	//})
	if err != nil {
		return nil, err
	}
	return i, nil
}

var _ I18ner = (*I18n)(nil)

func (i *I18n) Register(lang string, i18n interface{}) error {
	return i.adapter.Register(lang, i18n)
}

func (i *I18n) Update(lang, key string, i18n interface{}) error {
	return i.adapter.Update(lang, key, i18n)
}

func (i *I18n) SetDefault(lang string) error {
	return i.adapter.SetDefault(lang)
}

func (i *I18n) T(lang string, key string) (string, string, error) {
	return i.adapter.T(lang, key)
}

func (i *I18n) OnlyT(lang string, key string) string {
	return i.adapter.OnlyT(lang, key)
}
