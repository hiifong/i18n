package i18n

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"slices"
	"strings"

	"gopkg.in/ini.v1"
)

type INI struct {
	// Dir locale dir path, default: ./
	Dir string
	// Filename locale.en-US.ini, filename is locale, default: locale
	Filename string
	// DefaultLang default lang, default: en-US
	DefaultLang string
	// Languages is language list
	Languages []string
	// Ini *ini.FIle map
	Ini map[string]*ini.File
}

type Option func(*INI)

func WithDir(dir string) Option {
	return func(i *INI) {
		if dir != "" {
			i.Dir = dir
		}
	}
}

func WithFileName(name string) Option {
	return func(i *INI) {
		i.Filename = name
	}
}

func WithDefLang(lang string) Option {
	return func(i *INI) {
		if i.DefaultLang != "" {
			i.DefaultLang = lang
		}
	}
}

func WithLang(lang string) Option {
	return func(i *INI) {
		if lang == "" {
			return
		}
		if !slices.Contains(i.Languages, lang) {
			i.Languages = append(i.Languages, lang)
		}
	}
}

var _ Interface = (*INI)(nil)

func New(options ...Option) *INI {
	i := &INI{
		Dir:         "./",
		Filename:    "locale",
		DefaultLang: "en-US",
		Languages:   make([]string, 0, 20),
		Ini:         make(map[string]*ini.File),
	}
	for _, opt := range options {
		opt(i)
	}
	if !slices.Contains(i.Languages, i.DefaultLang) {
		i.Languages = append(i.Languages, i.DefaultLang)
	}
	return i
}

func (i *INI) Load() (err error) {
	for _, language := range i.Languages {
		var lerr error
		var path string
		if i.Filename == "" {
			path = filepath.Join(i.Dir, fmt.Sprintf("%s.ini", language))
		} else {
			path = filepath.Join(i.Dir, fmt.Sprintf("%s.%s.ini", i.Filename, language))
		}
		log.Printf("loading file %s", path)
		i.Ini[language], lerr = ini.Load(path)
		if lerr != nil {
			log.Printf("Error loading INI file %s: %s", path, lerr)
			err = errors.Join(lerr)
		}
	}
	return err
}

func (i *INI) Tr(lang, key string, values ...any) string {
	if _, ok := i.Ini[lang]; !ok {
		if _, ok = i.Ini[i.DefaultLang]; !ok {
			return ""
		} else {
			lang = i.DefaultLang
		}
	}
	if !strings.Contains(key, ".") {
		return fmt.Sprintf(i.Ini[lang].Section("").Key(key).String(), values...)
	}

	lastIndex := strings.LastIndex(key, ".")
	if lastIndex == -1 {
		return fmt.Sprintf(i.Ini[lang].Section("").Key(key).String(), values...)
	}

	section := strings.TrimSuffix(key[:lastIndex], ".")
	key = strings.TrimPrefix(key[lastIndex:], ".")
	return fmt.Sprintf(i.Ini[lang].Section(section).Key(key).String(), values...)
}
