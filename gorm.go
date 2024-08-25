package i18n

type gormI18n struct {
	ID      uint64 `gorm:"primary_key"`
	Lang    string
	Key     string
	Raw     string
	Created int64 `gorm:"autoCreateTime"`
	Updated int64 `gorm:"autoUpdateTime"`
	Deleted int64 `gorm:"index"`
}

func (i *gormI18n) TableName() string {
	return "i18n"
}

type gormOption func(*gormI18n)

func gormAdapterWithDefaultLang(lang string) gormOption {
	return func(i *gormI18n) {
		// TODO
	}
}

func gormAdapterWithLang(lang string, i18n interface{}) gormOption {
	return func(i *gormI18n) {
		// TODO
	}
}

func newGorm(dns string, options ...gormOption) (*gormI18n, error) {
	i18n := new(gormI18n)
	for _, option := range options {
		if option != nil {
			option(i18n)
		}
	}
	return i18n, nil
}

var _ adapter = (*gormI18n)(nil)

func (i *gormI18n) register(lang string, i18n interface{}) error {
	// TODO
	return nil
}

func (i *gormI18n) update(lang, key string, i18n interface{}) error {
	// TODO
	return nil
}

func (i *gormI18n) setDefault(lang string) error {
	// TODO
	return nil
}

func (i *gormI18n) t(lang string, key string) (string, string, error) {
	// TODO
	return "", "", nil
}

func (i *gormI18n) onlyT(lang string, key string) string {
	// TODO
	return ""
}
