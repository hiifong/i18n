package i18n

import "testing"

func TestI18n(t *testing.T) {
	i18n := New()
	zhCN := map[int]Language{
		1: {
			Code: 1,
			Raw:  "你好",
		},
		2: {
			Code: 2,
			Raw:  "世界",
		},
		3: {
			Code: 3,
			Raw:  "你好, 世界",
		},
	}
	enUS := map[int]Language{
		1: {
			Code: 1,
			Raw:  "Hello",
		},
		2: {
			Code: 2,
			Raw:  "World",
		},
		3: {
			Code: 3,
			Raw:  "Hello, World",
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
	code, msg, err := i18n.T("zh_CN", 1)
	if err != nil {
		t.Error(err)
	}
	t.Logf("code: %d, msg: %s", code, msg)

	code, msg, err = i18n.T("", 3)
	if err != nil {
		t.Error(err)
	}
	t.Logf("code: %d, msg: %s", code, msg)

	err = i18n.Update("zh_CN", 3, Language{
		Code: 3,
		Raw:  "Hello, World",
	})
	if err != nil {
		t.Error(err)
	}

	t.Logf("onlye t: %s", i18n.OnlyT("zh_CN", 3))

	code, msg, err = i18n.T("zh_CN", 4)
	if err != nil {
		t.Error(err)
	}
	t.Logf("zh_CN code 4: %d --> %s --> %v", code, msg, err)
}
