package i18n

import (
	"testing"
)

func TestI18n(t *testing.T) {
	zhCN := map[string]Language{
		"hello": {
			Key: "hello",
			Raw: "你好",
		},
		"world": {
			Key: "world",
			Raw: "世界",
		},
		"hello_world": {
			Key: "hello_world",
			Raw: "你好, 世界",
		},
	}
	enUS := map[string]Language{
		"hello": {
			Key: "hello",
			Raw: "Hello",
		},
		"world": {
			Key: "world",
			Raw: "World",
		},
		"hello_world": {
			Key: "hello_world",
			Raw: "Hello, World",
		},
	}
	i18n, _ := New(
		WithAdapter(Default),
		WithDefaultLang(EnUS),
		WithLang(ZhCN, zhCN),
	)
	err := i18n.Register(EnUS, enUS)
	if err != nil {
		t.Error(err)
	}
	key, msg, err := i18n.T(ZhCN, "hello")
	if err != nil {
		t.Error(err)
	}
	t.Logf("key: %s, msg: %s", key, msg)

	key, msg, err = i18n.T("", "hello_world")
	if err != nil {
		t.Error(err)
	}
	t.Logf("key: %s, msg: %s", key, msg)

	err = i18n.Update(ZhCN, "world", Language{
		Key: "world",
		Raw: "Hello, World",
	})
	if err != nil {
		t.Error(err)
	}

	key, msg, err = i18n.T(ZhCN, "world")
	if err != nil {
		t.Error(err)
	}
	t.Logf("key: %s, msg: %s", key, msg)

	err = i18n.Update(ZhCN, "hello", nil)
	if err != nil {
		t.Error(err)
	}

	err = i18n.Update(ZhCN, "unknown", Language{})
	if err != nil {
		t.Error(err)
	}

	t.Logf("onlye t: %s", i18n.OnlyT(ZhCN, "hello_world"))

	key, msg, err = i18n.T(ZhHK, "hello")
	if err != nil {
		t.Error(err)
	}
	t.Logf("zh_HK key: %s --> msg: %s --> err: %v", key, msg, err)
}
