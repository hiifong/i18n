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

type adapter interface {
	// Register 注册新的语言
	register(lang string, i18n interface{}) error
	// Update 更新翻译, 如果存在翻译则更新，否则添加翻译
	update(lang, key string, i18n interface{}) error
	// SetDefault 设置默认语言
	setDefault(lang string) error
	// T 获取翻译
	t(lang, key string) (string, string, error)
	// OnlyT 仅获取翻译
	onlyT(lang string, key string) string
}

type Type string

const (
	Default Type = "default"
	Redis   Type = "redis"
	Ent     Type = "ent"
	Gorm    Type = "gorm"

	i18nKey string = "i18n"
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
	adapter adapter
}

type Option func(*I18n)

// WithAdapter dns like:
// default: nil
// redis: redis://user:password@localhost:6379/0?protocol=3
// mysql:
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
		case Redis:
			i.opts = append(i.opts, redisAdapterWithDefaultLang(lang))
		case Ent:
			i.opts = append(i.opts, entAdapterWithDefaultLang(lang))
		case Gorm:
			i.opts = append(i.opts, gormAdapterWithDefaultLang(lang))
		case Default:
			i.opts = append(i.opts, defaultAdapterWithDefaultLang(lang))
		default:
			i.opts = append(i.opts, defaultAdapterWithDefaultLang(lang))
		}
	}
}

func WithLang(lang string, i18n interface{}) Option {
	return func(i *I18n) {
		switch i.t {
		case Redis:
			i.opts = append(i.opts, redisAdapterWithLang(lang, i18n))
		case Ent:
			i.opts = append(i.opts, entAdapterWithLang(lang, i18n))
		case Gorm:
			i.opts = append(i.opts, gormAdapterWithLang(lang, i18n))
		case Default:
			i.opts = append(i.opts, defaultAdapterWithLang(lang, i18n))
		default:
			i.opts = append(i.opts, defaultAdapterWithLang(lang, i18n))
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
	case Redis:
		var opts []redisOption
		for _, opt := range i.opts {
			opts = append(opts, opt.(redisOption))
		}
		i.adapter, err = newRedis(i.dns, opts...)
	case Ent:
		var opts []entOption
		for _, opt := range i.opts {
			opts = append(opts, opt.(entOption))
		}
		i.adapter, err = newEnt(i.dns, opts...)
	case Gorm:
		var opts []gormOption
		for _, opt := range i.opts {
			opts = append(opts, opt.(gormOption))
		}
		i.adapter, err = newGorm(i.dns, opts...)
	default:
		var opts []defaultOption
		for _, opt := range i.opts {
			opts = append(opts, opt.(defaultOption))
		}
		i.adapter, err = newDefault(opts...)
	}
	//})
	if err != nil {
		return nil, err
	}
	return i, nil
}

var _ I18ner = (*I18n)(nil)

func (i *I18n) Register(lang string, i18n interface{}) error {
	return i.adapter.register(lang, i18n)
}

func (i *I18n) Update(lang, key string, i18n interface{}) error {
	return i.adapter.update(lang, key, i18n)
}

func (i *I18n) SetDefault(lang string) error {
	return i.adapter.setDefault(lang)
}

func (i *I18n) T(lang string, key string) (string, string, error) {
	return i.adapter.t(lang, key)
}

func (i *I18n) OnlyT(lang string, key string) string {
	return i.adapter.onlyT(lang, key)
}
