# i18n
golang i18n

## install

```shell
go get github.com/hiifong/i18n@v1.0.2
```

## i18n interface

```go
// I18ner 国际化接口
type I18ner interface {
	// Register 注册新的语言
	Register(lang string, i18n interface{}) error

	// Update 更新翻译, 如果存在翻译则更新，否则添加翻译 
	Update(lang string, code int, t interface{}) error

	// SetDefault 设置默认语言
	SetDefault(lang string) error

	// T 获取翻译
	T(lang string, code int) (int, string, error)

	// OnlyT 仅获取翻译
	OnlyT(lang string, code int) string
}
```

## Example

```go
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

	t.Logf("onlye t: %s", i18n.OnlyT("", 3))
}
```