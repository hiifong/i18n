# i18n
golang i18n

## install

```shell
go get github.com/hiifong/i18n@v1.3.1
```

## i18n interface

```go
// I18ner 国际化接口
type I18ner interface {
        // Register 注册新的语言
        Register(lang string, i18n interface{}) error
        // Update 更新翻译, 如果存在翻译则更新，否则添加翻译
        Update(lang, key string, i18n interface{}) error
        // SetDefault 设置默认语言
        SetDefault(lang string) error
        // T 获取翻译
        T(lang, key string) (string, string, error)
        // OnlyT 仅获取翻译
        OnlyT(lang string, key string) string
}

```

## Example

```go
package main

import (
	"fmt"

	"github.com/hiifong/i18n"
)

func main() {
	zhCN := map[string]i18n.Language{
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
	enUS := map[string]i18n.Language{
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
	i, _ := i18n.New(
		i18n.WithAdapter(i18n.Default),
		i18n.WithDefaultLang(i18n.EnUS),
		i18n.WithLang(i18n.ZhCN, zhCN),
	)
	err := i.Register(i18n.ZhCN, zhCN)
	if err != nil {
		fmt.Println(err)
	}
	err = i.Register(i18n.EnUS, enUS)
	if err != nil {
		fmt.Println(err)
	}
	code, msg, err := i.T(i18n.ZhCN, "hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("key: %s, msg: %s\n", code, msg)

	code, msg, err = i.T("", "hello_world")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("key: %s, msg: %s\n", code, msg)

	err = i.Update(i18n.EnUS, "world", i18n.Language{
		Key: "world",
		Raw: "Hello, World",
	})
	if err != nil {
		fmt.Println(err)
	}
	key, msg, err := i.T(i18n.EnUS, "world")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("key: %s, msg: %s\n", key, msg)

	err = i.Update(i18n.ZhCN, "hello", nil)
	if err != nil {
		fmt.Println(err)
	}

	err = i.Update(i18n.ZhCN, "unknown", i18n.Language{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("onlye t: %s\n", i.OnlyT(i18n.ZhCN, "hello_world"))

	key, msg, err = i.T(i18n.ZhHK, "hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("zh_HK key: %s --> msg: %s --> err:%v\n", key, msg, err)
}

// Output:
language zh-CN is already registered
key: hello, msg: 你好
key: hello_world, msg: Hello, World
key: world, msg: Hello, World
this zh-CN is not support
this zh-CN key is not match
onlye t: 你好, 世界
zh_HK key: hello --> msg: Hello --> err:<nil>
```

## 附件 - 国际化开发的各国语言标识
语言标识|国家地区  
---|:---  
zh_CN  |  简体中文(中国)  
zh_TW  |  繁体中文(台湾地区)  
zh_HK  |  繁体中文(香港)  
en_HK  |  英语(香港)  
en_US  |  英语(美国)  
en_GB  |  英语(英国)  
en_WW  |  英语(全球)  
en_CA  |  英语(加拿大)  
en_AU  |  英语(澳大利亚)  
en_IE  |  英语(爱尔兰)  
en_FI  |  英语(芬兰)  
fi_FI  |  芬兰语(芬兰)  
en_DK  |  英语(丹麦)  
da_DK  |  丹麦语(丹麦)  
en_IL  |  英语(以色列)  
he_IL  |  希伯来语(以色列)  
en_ZA  |  英语(南非)  
en_IN  |  英语(印度)  
en_NO  |  英语(挪威)  
en_SG  |  英语(新加坡)  
en_NZ  |  英语(新西兰)  
en_ID  |  英语(印度尼西亚)  
en_PH  |  英语(菲律宾)  
en_TH  |  英语(泰国)  
en_MY  |  英语(马来西亚)  
en_XA  |  英语(阿拉伯)  
ko_KR  |  韩文(韩国)  
ja_JP  |  日语(日本)  
nl_NL  |  荷兰语(荷兰)  
nl_BE  |  荷兰语(比利时)  
pt_PT  |  葡萄牙语(葡萄牙)  
pt_BR  |  葡萄牙语(巴西)  
fr_FR  |  法语(法国)  
fr_LU  |  法语(卢森堡)  
fr_CH  |  法语(瑞士)  
fr_BE  |  法语(比利时)  
fr_CA  |  法语(加拿大)  
es_LA  |  西班牙语(拉丁美洲)  
es_ES  |  西班牙语(西班牙)  
es_AR  |  西班牙语(阿根廷)  
es_US  |  西班牙语(美国)  
es_MX  |  西班牙语(墨西哥)  
es_CO  |  西班牙语(哥伦比亚)  
es_PR  |  西班牙语(波多黎各)  
de_DE  |  德语(德国)  
de_AT  |  德语(奥地利)  
de_CH  |  德语(瑞士)  
ru_RU  |  俄语(俄罗斯)  
it_IT  |  意大利语(意大利)  
el_GR  |  希腊语(希腊)  
no_NO  |  挪威语(挪威)  
hu_HU  |  匈牙利语(匈牙利)  
tr_TR  |  土耳其语(土耳其)  
cs_CZ  |  捷克语(捷克共和国)  
sl_SL  |  斯洛文尼亚语   
pl_PL  |  波兰语(波兰)  
sv_SE  |  瑞典语(瑞典)  
es_CL  |  西班牙语 (智利)  