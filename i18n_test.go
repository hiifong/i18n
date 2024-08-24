package i18n

import "testing"

func TestI18n(t *testing.T) {
	i18n := New()
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
	err := i18n.Register("zh_CN", zhCN)
	if err != nil {
		t.Error(err)
	}
	err = i18n.Register("en_US", enUS)
	if err != nil {
		t.Error(err)
	}
	err = i18n.SetDefault("en_US")
	if err != nil {
		t.Error(err)
	}
	code, msg, err := i18n.T("zh_CN", "hello")
	if err != nil {
		t.Error(err)
	}
	t.Logf("key: %s, msg: %s", code, msg)

	code, msg, err = i18n.T("", "hello_world")
	if err != nil {
		t.Error(err)
	}
	t.Logf("key: %s, msg: %s", code, msg)

	err = i18n.Update("zh_CN", "world", Language{
		Key: "hello_world",
		Raw: "Hello, World",
	})
	if err != nil {
		t.Error(err)
	}

	err = i18n.Update("zh_CN", "hello", nil)
	if err != nil {
		t.Error(err)
	}

	err = i18n.Update("zh_CN", "unknown", Language{})
	if err != nil {
		t.Error(err)
	}

	t.Logf("onlye t: %s", i18n.OnlyT("zh_CN", "hello_world"))

	code, msg, err = i18n.T("zh_CN", "unknown")
	if err != nil {
		t.Error(err)
	}
	t.Logf("zh_CN code 4: %s --> %s --> %v", code, msg, err)
}
