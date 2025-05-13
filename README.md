# i18n
golang i18n

## install

```shell
go get -u github.com/hiifong/i18n
```

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/hiifong/i18n"
)

func main() {
	i18n.SetDefault(
		i18n.New(
			i18n.WithDir("/Users/hiifong/Desktop/Golang/i18n/example"),
			i18n.WithLang("zh-CN"),
		),
	)
	err := i18n.Load()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(i18n.Tr("en-US", "hello"))
	fmt.Println(i18n.Tr("zh-CN", "hello_to", "hiifong"))
	fmt.Println(i18n.Tr("en-US", "db.host"))
	fmt.Println(i18n.Tr("en-US", "db.port"))
	fmt.Println(i18n.Tr("en-US", "cache.redis.host"))
	fmt.Println(i18n.Tr("en-US", "cache.redis.port"))
	fmt.Println(i18n.Tr("en-US", "db.host"))
	fmt.Println(i18n.Tr("zh-CN", "db.port"))
	fmt.Println(i18n.Tr("zh-CN", "cache.redis.host"))
	fmt.Println(i18n.Tr("zh-CN", "cache.redis.port"))
	fmt.Println(i18n.Tr("zh-HK", "cache.redis.host"))
}

// Output:
hello
你好, hiifong
database host
database port
redis cache host
redis cache port
database host
数据库端口
redis缓存主机
redis缓存端口
redis cache host
```

## 附件 - 国际化开发的各国语言标识
语言标识|国家地区  
---|:---  
zh_CN  |  简体中文(中国)  
zh_HK  |  繁体中文(中国香港)
zh_MO  |  繁体中文(中国澳门)
zh_TW  |  繁体中文(中国台湾)  
en_HK  |  英语(中国香港)  
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