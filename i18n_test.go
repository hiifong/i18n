package i18n

import (
	"testing"
)

func TestI18n(t *testing.T) {
	zhCN := []Language{
		{
			Key: "hello",
			Raw: "你好",
		},
		{
			Key: "world",
			Raw: "世界",
		},
		{
			Key: "hello_world",
			Raw: "你好, 世界",
		},
	}
	enUS := []Language{
		{
			Key: "hello",
			Raw: "Hello",
		},
		{
			Key: "world",
			Raw: "World",
		},
		{
			Key: "hello_world",
			Raw: "Hello, World",
		},
	}
	i18n, _ := New(
		WithAdapter(Redis, "redis://192.168.5.104:6379/0?protocol=3"),
		WithDefaultLang(EnUS),
		WithLang(ZhCN, zhCN),
	)
	err := i18n.Register(EnUS, enUS)
	if err != nil {
		t.Error(err)
	}
	key, text, err := i18n.T(ZhCN, "hello")
	if err != nil {
		t.Error(err)
	}
	t.Logf("key: %s, text: %s", key, text)

	key, text, err = i18n.T("", "hello_world")
	if err != nil {
		t.Error(err)
	}
	t.Logf("key: %s, text: %s", key, text)
	key, text, err = i18n.T(ZhCN, "world")
	if err != nil {
		t.Error(err)
	}
	t.Logf("key: %s, text: %s\n", key, text)

	err = i18n.Update(ZhCN, "world", Language{
		Key: "world",
		Raw: "Hello, World",
	})
	if err != nil {
		t.Error(err)
	}

	key, text, err = i18n.T(ZhCN, "world")
	if err != nil {
		t.Error(err)
	}
	t.Logf("key: %s, text: %s\n", key, text)

	err = i18n.Update(ZhCN, "hello", nil)
	if err != nil {
		t.Error(err)
	}

	err = i18n.Update(ZhCN, "unknown", Language{})
	if err != nil {
		t.Error(err)
	}

	t.Logf("onlye t: %s\n", i18n.OnlyT(ZhCN, "hello_world"))

	key, text, err = i18n.T(ZhHK, "hello")
	if err != nil {
		t.Error(err)
	}
	t.Logf("zh_HK key: %s --> text: %s --> err: %v\n", key, text, err)
}
