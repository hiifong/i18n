package i18n

import (
	"context"
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
)

type redisI18n struct {
	mu    sync.RWMutex
	first string
	i18n  *redis.Client
}

type redisOption func(*redisI18n)

func redisAdapterWithDefaultLang(lang string) redisOption {
	return func(i *redisI18n) {
		i.first = fmt.Sprintf("%s:%s", i18nKey, lang)
	}
}

func redisAdapterWithLang(lang string, i18n interface{}) redisOption {
	return func(i *redisI18n) {
		i.mu.Lock()
		defer i.mu.Unlock()
		ll := i18n.([]Language)
		lm := make(map[string]string)
		for _, l := range ll {
			lm[l.Key] = l.Raw
		}
		i.i18n.HSet(context.Background(), fmt.Sprintf("%s:%s", i18nKey, lang), lm)
	}
}

// newRedis dns like: redis://user:password@localhost:6379/0?protocol=3
func newRedis(dns string, options ...redisOption) (*redisI18n, error) {
	i := new(redisI18n)
	if dns == "" {
		return i, fmt.Errorf("dns can't be empty")
	}
	opts, err := redis.ParseURL(dns)
	if err != nil {
		return nil, fmt.Errorf("redis dns parse url err: %w", err)
	}
	i.i18n = redis.NewClient(opts)
	for _, option := range options {
		if option != nil {
			option(i)
		}
	}
	return i, nil
}

var _ adapter = (*redisI18n)(nil)

func (i *redisI18n) register(lang string, i18n interface{}) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	if lang == "" {
		return fmt.Errorf("lang can't be empty")
	}
	if i.i18n.Exists(context.Background(), fmt.Sprintf("%s:%s", i18nKey, lang)).Val() > 0 {
		return fmt.Errorf("i18n %s is already registered", lang)
	}
	ll, ok := i18n.([]Language)
	if !ok {
		return fmt.Errorf("i18n %s is invalid", lang)
	}
	lm := make(map[string]string)
	for _, l := range ll {
		lm[l.Key] = l.Raw
	}
	err := i.i18n.HSet(context.Background(), fmt.Sprintf("%s:%s", i18nKey, lang), lm).Err()
	if err != nil {
		return fmt.Errorf("register language %s, i18n: %v, err: %w", lang, i18n, err)
	}
	return nil
}

func (i *redisI18n) update(lang, key string, i18n interface{}) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	if lang == "" {
		return fmt.Errorf("language can't be empty")
	}
	if i.i18n.Exists(context.Background(), fmt.Sprintf("%s:%s", i18nKey, lang)).Val() <= 0 {
		return fmt.Errorf("language %s is not registered", lang)
	}
	l, ok := i18n.(Language)
	if !ok {
		return fmt.Errorf("language %s is invalid", lang)
	}
	if key != l.Key {
		return fmt.Errorf("this %s key is not match", lang)
	}
	err := i.i18n.HSet(context.Background(), fmt.Sprintf("%s:%s", i18nKey, lang), l.Key, l.Raw).Err()
	if err != nil {
		return fmt.Errorf("update language %s, i18n: %v, err: %w", lang, i18n, err)
	}
	return nil
}

func (i *redisI18n) setDefault(lang string) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	if lang == "" {
		return fmt.Errorf("lang can't be empty")
	}
	if i.i18n.Exists(context.Background(), fmt.Sprintf("%s:%s", i18nKey, lang)).Val() <= 0 {
		return fmt.Errorf("language %s is already registered", lang)
	}
	i.first = lang
	return nil
}

func (i *redisI18n) t(lang string, key string) (string, string, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	if i.first == "" {
		return key, "", fmt.Errorf("default language can't be empty")
	}
	if key == "" {
		return key, "", fmt.Errorf("key can't be empty")
	}
	t := func(lang string, key string) (string, string, error) {
		ok := i.i18n.Exists(context.Background(), fmt.Sprintf("%s:%s", i18nKey, lang)).Val() > 0
		if !ok {
			ok = i.i18n.Exists(context.Background(), i.first).Val() > 0
		}
		if !ok {
			return key, "", fmt.Errorf("language %s is not registered", lang)
		}
		text := ""
		if i.i18n.HExists(context.Background(), fmt.Sprintf("%s:%s", i18nKey, lang), key).Val() {
			text = i.i18n.HGet(context.Background(), fmt.Sprintf("%s:%s", i18nKey, lang), key).Val()
			if text == "" {
				return key, "", fmt.Errorf("i18n is empty")
			}
			return key, text, nil
		}
		if i.i18n.HExists(context.Background(), i.first, key).Val() {
			text = i.i18n.HGet(context.Background(), i.first, key).Val()
			if text == "" {
				return key, "", fmt.Errorf("i18n is empty")
			}
			return key, text, nil
		}
		return key, text, fmt.Errorf("not found")
	}

	if lang == "" {
		return t(i.first, key)
	}
	return t(lang, key)
}

func (i *redisI18n) onlyT(lang string, key string) string {
	_, text, err := i.t(lang, key)
	if err != nil {
		return ""
	}
	return text
}
