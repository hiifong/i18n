package i18n

type Lang string

var (
	languages = []Lang{
		ZhCN, ZhHK, ZhMO, ZhTW, EnHK, EnUS, EnGB, EnWW, EnCA, EnAU,
		EnIE, EnFI, FiFI, EnDK, DaDK, EnIL, HeIL, EnZA, EnIN, EnNO,
		EnSG, EnNZ, EnID, EnPH, EnTH, EnMY, EnXA, KoKR, JaJP, NlNL,
		NlBE, PtPT, PtBR, FrFR, FrLU, FrCH, FrBE, FrCA, EsLA, EsES,
		EsAR, EsUS, EsMX, EsCO, EsPR, DeDE, DeAT, DeCH, RuRU, ItIT,
		ElGR, NoNO, HuHU, TrTR, CsCZ, SlSL, PlPL, SvSE, EsCL,
	}
	AllLanguages = func() []Lang {
		return languages
	}()
)

const (
	ZhCN Lang = "zh-CN" // 简体中文(中国)
	ZhHK Lang = "zh-HK" // 繁体中文(中国香港)
	ZhMO Lang = "zh-MO" // 繁体中文(中国澳门)
	ZhTW Lang = "zh-TW" // 繁体中文(中国台湾)
	EnHK Lang = "en-HK" // 英语(中国香港)
	EnUS Lang = "en-US" // 英语(美国)
	EnGB Lang = "en-GB" // 英语(英国)
	EnWW Lang = "en-WW" // 英语(全球)
	EnCA Lang = "en-CA" // 英语(加拿大)
	EnAU Lang = "en-AU" // 英语(澳大利亚)
	EnIE Lang = "en-IE" // 英语(爱尔兰)
	EnFI Lang = "en-FI" // 英语(芬兰)
	FiFI Lang = "fi-FI" // 芬兰语(芬兰)
	EnDK Lang = "en-DK" // 英语(丹麦)
	DaDK Lang = "da-DK" // 丹麦语(丹麦)
	EnIL Lang = "en-IL" // 英语(以色列)
	HeIL Lang = "he-IL" // 希伯来语(以色列)
	EnZA Lang = "en-ZA" // 英语(南非)
	EnIN Lang = "en-IN" // 英语(印度)
	EnNO Lang = "en-NO" // 英语(挪威)
	EnSG Lang = "en-SG" // 英语(新加坡)
	EnNZ Lang = "en-NZ" // 英语(新西兰)
	EnID Lang = "en-ID" // 英语(印度尼西亚)
	EnPH Lang = "en-PH" // 英语(菲律宾)
	EnTH Lang = "en-TH" // 英语(泰国)
	EnMY Lang = "en-MY" // 英语(马来西亚)
	EnXA Lang = "en-XA" // 英语(阿拉伯)
	KoKR Lang = "ko-KR" // 韩文(韩国)
	JaJP Lang = "ja-JP" // 日语(日本)
	NlNL Lang = "nl-NL" // 荷兰语(荷兰)
	NlBE Lang = "nl-BE" // 荷兰语(比利时)
	PtPT Lang = "pt-PT" // 葡萄牙语(葡萄牙)
	PtBR Lang = "pt-BR" // 葡萄牙语(巴西)
	FrFR Lang = "fr-FR" // 法语(法国)
	FrLU Lang = "fr-LU" // 法语(卢森堡)
	FrCH Lang = "fr-CH" // 法语(瑞士)
	FrBE Lang = "fr-BE" // 法语(比利时)
	FrCA Lang = "fr-CA" // 法语(加拿大)
	EsLA Lang = "es-LA" // 西班牙语(拉丁美洲)
	EsES Lang = "es-ES" // 西班牙语(西班牙)
	EsAR Lang = "es-AR" // 西班牙语(阿根廷)
	EsUS Lang = "es-US" // 西班牙语(美国)
	EsMX Lang = "es-MX" // 西班牙语(墨西哥)
	EsCO Lang = "es-CO" // 西班牙语(哥伦比亚)
	EsPR Lang = "es-PR" // 西班牙语(波多黎各)
	DeDE Lang = "de-DE" // 德语(德国)
	DeAT Lang = "de-AT" // 德语(奥地利)
	DeCH Lang = "de-CH" // 德语(瑞士)
	RuRU Lang = "ru-RU" // 俄语(俄罗斯)
	ItIT Lang = "it-IT" // 意大利语(意大利)
	ElGR Lang = "el-GR" // 希腊语(希腊)
	NoNO Lang = "no-NO" // 挪威语(挪威)
	HuHU Lang = "hu-HU" // 匈牙利语(匈牙利)
	TrTR Lang = "tr-TR" // 土耳其语(土耳其)
	CsCZ Lang = "cs-CZ" // 捷克语(捷克共和国)
	SlSL Lang = "sl-SL" // 斯洛文尼亚语
	PlPL Lang = "pl-PL" // 波兰语(波兰)
	SvSE Lang = "sv-SE" // 瑞典语(瑞典)
	EsCL Lang = "es-CL" // 西班牙语 (智利)
)
