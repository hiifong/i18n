package i18n

type GormI18n struct {
	ID      uint64 `gorm:"primary_key"`
	Lang    string
	Key     string
	Raw     string
	Created int64 `gorm:"autoCreateTime"`
	Updated int64 `gorm:"autoUpdateTime"`
	Deleted int64 `gorm:"index"`
}

func (i *GormI18n) TableName() string {
	return "i18n"
}

type GormOption func(n *GormI18n)

func GormAdapterWithDefaultLang(lang string) GormOption {
	return func(i *GormI18n) {
		// TODO
	}
}

func GormAdapterWithLang(lang string, i18n interface{}) GormOption {
	return func(i *GormI18n) {
		// TODO
	}
}

func NewGorm(dns string, options ...GormOption) (*GormI18n, error) {
	i18n := new(GormI18n)
	for _, option := range options {
		if option != nil {
			option(i18n)
		}
	}
	return i18n, nil
}

var _ I18ner = (*GormI18n)(nil)

func (i *GormI18n) Register(lang string, i18n interface{}) error {
	// TODO
	return nil
}

func (i *GormI18n) Update(lang, key string, i18n interface{}) error {
	// TODO
	return nil
}

func (i *GormI18n) SetDefault(lang string) error {
	// TODO
	return nil
}

func (i *GormI18n) T(lang string, key string) (string, string, error) {
	// TODO
	return "", "", nil
}

func (i *GormI18n) OnlyT(lang string, key string) string {
	// TODO
	return ""
}