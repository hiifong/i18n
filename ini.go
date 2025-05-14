package i18n

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	"gopkg.in/ini.v1"
)

type INI struct {
	mu sync.RWMutex
	// dir locale dir path, default: ./
	dir string
	// fs locale filesystem
	fs http.FileSystem
	// auto automatically detect locale and load, filename format must be {filename}.{lang}.ini
	auto bool
	// filename locale.en-US.ini, filename is locale, default: locale
	filename string
	// defaultLang default lang, default: en-US
	defaultLang Lang
	// languages is language list
	languages []Lang
	// iniFileMap *ini.FIle map
	iniFileMap map[Lang]*ini.File
}

type Option func(*INI)

func WithDir(dir string) Option {
	return func(i *INI) {
		if dir != "" {
			i.dir = dir
		}
	}
}

func WithFS(fs http.FileSystem) Option {
	return func(i *INI) {
		if fs != nil {
			i.fs = fs
		}
	}
}

func WithAuto() Option {
	return func(i *INI) {
		i.auto = true
	}
}

func WithFileName(name string) Option {
	return func(i *INI) {
		i.filename = name
	}
}

func WithDefLang(lang Lang) Option {
	return func(i *INI) {
		if i.defaultLang != "" {
			i.defaultLang = lang
		}
	}
}

func WithLang(lang Lang) Option {
	return func(i *INI) {
		if lang == "" {
			return
		}
		if !slices.Contains(i.languages, lang) {
			i.languages = append(i.languages, lang)
		}
	}
}

var _ Interface = (*INI)(nil)

func New(options ...Option) *INI {
	i := &INI{
		dir:         "./",
		filename:    "locale",
		defaultLang: "en-US",
		languages:   make([]Lang, 0, 20),
		iniFileMap:  make(map[Lang]*ini.File),
	}
	for _, opt := range options {
		opt(i)
	}
	if !slices.Contains(i.languages, i.defaultLang) {
		i.languages = append(i.languages, i.defaultLang)
	}
	return i
}

func (i *INI) Load() error {
	i.mu.Lock()
	defer i.mu.Unlock()
	if i.auto {
		return i.autoLoad()
	}
	if i.fs == nil {
		return i.loadFromDir()
	}
	return i.loadFromFS()
}

func (i *INI) autoLoad() error {
	for _, lang := range AllLanguages {
		var path string
		if i.filename == "" {
			path = filepath.Join(i.dir, fmt.Sprintf("%s.ini", lang))
		} else {
			path = filepath.Join(i.dir, fmt.Sprintf("%s.%s.ini", i.filename, lang))
		}
		log.Printf("load locale file %s", path)
		if i.fs != nil {
			if file, err := i.fs.Open(path); err == nil {
				defer file.Close()
				i.iniFileMap[lang], err = ini.Load(file)
				if err != nil {
					log.Printf("fail to load locale file %s err %v", path, err)
					continue
				}
				if !slices.Contains(i.languages, lang) {
					i.languages = append(i.languages, lang)
				}
			} else {
				log.Printf("fail to load locale file %s err %v", path, err)
			}
		} else {
			var err error
			i.iniFileMap[lang], err = ini.Load(path)
			if err != nil {
				log.Printf("fail to load locale file %s err %v", path, err)
				continue
			}
			if !slices.Contains(i.languages, lang) {
				i.languages = append(i.languages, lang)
			}
		}
	}
	return nil
}

func (i *INI) loadFromDir() error {
	var err error
	for _, language := range i.languages {
		var lerr error
		var path string
		if i.filename == "" {
			path = filepath.Join(i.dir, fmt.Sprintf("%s.ini", language))
		} else {
			path = filepath.Join(i.dir, fmt.Sprintf("%s.%s.ini", i.filename, language))
		}
		log.Printf("loading file %s", path)
		i.iniFileMap[language], lerr = ini.Load(path)
		if lerr != nil {
			log.Printf("Error loading INI file %s: %s", path, lerr)
			err = errors.Join(lerr)
		}
	}
	return err
}

func (i *INI) loadFromFS() error {
	var err error
	for _, language := range i.languages {
		var lerr error
		var path string
		if i.filename == "" {
			path = filepath.Join(i.dir, fmt.Sprintf("%s.ini", language))
		} else {
			path = filepath.Join(i.dir, fmt.Sprintf("%s.%s.ini", i.filename, language))
		}
		log.Printf("loading file %s", path)
		file, lerr := i.fs.Open(path)
		if lerr != nil {
			log.Printf("Error loading INI file %s from fs: %s", path, lerr)
			err = errors.Join(lerr)
		}
		defer file.Close()
		i.iniFileMap[language], lerr = ini.Load(file)
		if lerr != nil {
			log.Printf("Error loading INI file %s: %s", path, lerr)
			err = errors.Join(lerr)
		}
	}
	return err
}

func (i *INI) Default() Lang {
	return i.defaultLang
}

func (i *INI) SetDefault(lang Lang) {
	i.defaultLang = lang
}

func (i *INI) Tr(lang Lang, key string, values ...any) template.HTML {
	i.mu.RLock()
	defer i.mu.RUnlock()
	if file, ok := i.iniFileMap[lang]; !ok || file == nil {
		if file, ok := i.iniFileMap[i.defaultLang]; !ok || file == nil {
			return ""
		} else {
			lang = i.defaultLang
		}
	}
	if file, ok := i.iniFileMap[lang]; !ok || file == nil {
		return ""
	}

	if !strings.Contains(key, ".") {
		return template.HTML(fmt.Sprintf(i.iniFileMap[lang].Section("").Key(key).String(), values...))
	}

	lastIndex := strings.LastIndex(key, ".")

	section := strings.TrimSuffix(key[:lastIndex], ".")
	key = strings.TrimPrefix(key[lastIndex:], ".")
	return template.HTML(fmt.Sprintf(i.iniFileMap[lang].Section(section).Key(key).String(), values...))
}

func (i *INI) Languages() []Lang {
	return i.languages
}
