package main

type card struct {
	character string
	pinyin string
	meaning string
}

var cards = []card {
	{
		character: "的",
		pinyin: "de",
		meaning: "possessive particle",
	},
	{
		character: "一",
		pinyin: "yi1",
		meaning: "one",
	},
	{
		character: "是",
		pinyin: "shi4",
		meaning: "to be",
	},
	{
		character: "不",
		pinyin: "bu4",
		meaning: "not, negative",
	},
	{
		character: "了",
		pinyin: "le",
		meaning: "verb particle marking past tense/completion",
	},
	{
		character: "人",
		pinyin: "ren2",
		meaning: "person",
	},
	{
		character: "我",
		pinyin: "wo3",
		meaning: "I",
	},
	{
		character: "在",
		pinyin: "zai4",
		meaning: "at, located at",
	},
	{
		character: "有",
		pinyin: "you3",
		meaning: "have, there is",
	},
	{
		character: "他",
		pinyin: "ta1",
		meaning: "he",
	},
	{
		character: "這",
		pinyin: "zhe4",
		meaning: "this",
	},
	{
		character: "中",
		pinyin: "zhong1",
		meaning: "middle",
	},
	{
		character: "大",
		pinyin: "da4",
		meaning: "big",
	},
	{
		character: "來",
		pinyin: "lai2",
		meaning: "come",
	},
	{
		character: "上",
		pinyin: "shang4",
		meaning: "above, on, over/last, previous",
	},
	{
		character: "國",
		pinyin: "guo2",
		meaning: "country",
	},
	{
		character: "個",
		pinyin: "ge4",
		meaning: "general quantifier",
	},
	{
		character: "到",
		pinyin: "dao4",
		meaning: "to, towards, until, arrive",
	},
	{
		character: "說",
		pinyin: "shuo1",
		meaning: "speak, say",
	},
	{
		character: "們",
		pinyin: "men",
		meaning: "plural noun referring to people",
	},
	{
		character: "為",
		pinyin: "wei4",
		meaning: "in order to / for the purpose of",
	},
	{
		character: "子",
		pinyin: "zi3",
		meaning: "child",
	},
	{
		character: "和",
		pinyin: "he2",
		meaning: "together, with",
	},
	{
		character: "你",
		pinyin: "ni3",
		meaning: "you",
	},
	{
		character: "地",
		pinyin: "di4",
		meaning: "earth, ground, place",
	},
	{
		character: "出",
		pinyin: "chu1",
		meaning: "go out, come out",
	},
	{
		character: "道",
		pinyin: "dao4",
		meaning: "way, path",
	},
	{
		character: "也",
		pinyin: "ye3",
		meaning: "also, as well",
	},
	{
		character: "時",
		pinyin: "shi2",
		meaning: "time period, season",
	},
	{
		character: "年",
		pinyin: "nian2",
		meaning: "year",
	},
	{
		character: "得",
		pinyin: "de",
		meaning: "to have to / adverbial particle",
	},
	{
		character: "就",
		pinyin: "jiu4",
		meaning: "just, simply, right away",
	},
	{
		character: "那",
		pinyin: "na4",
		meaning: "that",
	},
	{
		character: "要",
		pinyin: "yao4",
		meaning: "want, will do, important, need",
	},
	{
		character: "下",
		pinyin: "xia4",
		meaning: "below, down, under, next",
	},
	{
		character: "以",
		pinyin: "yi3",
		meaning: "because of, in order to, according to",
	},
	{
		character: "生",
		pinyin: "sheng1",
		meaning: "give birth, life",
	},
	{
		character: "會",
		pinyin: "hui4",
		meaning: "can, able, meet, party, society",
	},
	{
		character: "自",
		pinyin: "zi4",
		meaning: "from, self",
	},
	{
		character: "着",
		pinyin: "zhe",
		meaning: "verb particle marking a progressive aspect (e.g. -ing)",
	},
	{
		character: "去",
		pinyin: "qu4",
		meaning: "go, leave, depart",
	},
	{
		character: "之",
		pinyin: "zhi1",
		meaning: "subordinator, go, leave for",
	},
	{
		character: "過",
		pinyin: "guo4",
		meaning: "pass, cross, exceed",
	},
	{
		character: "家",
		pinyin: "jia1",
		meaning: "home, family",
	},
	{
		character: "學",
		pinyin: "xue2",
		meaning: "study, learn",
	},
	{
		character: "對",
		pinyin: "dui4",
		meaning: "correct, mutual, pair",
	},
	{
		character: "可",
		pinyin: "ke3",
		meaning: "~able, may, can",
	},
	{
		character: "她",
		pinyin: "ta1",
		meaning: "she",
	},
	{
		character: "裡",
		pinyin: "li3",
		meaning: "neighbourhood, half kilometer",
	},
	{
		character: "后",
		pinyin: "hou4",
		meaning: "queen",
	},
	{
		character: "小",
		pinyin: "xiao3",
		meaning: "small",
	},
	{
		character: "麼",
		pinyin: "me/mo",
		meaning: "interrogative suffix",
	},
	{
		character: "心",
		pinyin: "xin1",
		meaning: "heart",
	},
	{
		character: "多",
		pinyin: "duo1",
		meaning: "many, much, more",
	},
	{
		character: "天",
		pinyin: "tian1",
		meaning: "sky, heaven, god, day",
	},
	{
		character: "而",
		pinyin: "er2",
		meaning: "and, furthermore",
	},
	{
		character: "能",
		pinyin: "neng2",
		meaning: "can, be able",
	},
	{
		character: "好",
		pinyin: "hao3",
		meaning: "good",
	},
	{
		character: "都",
		pinyin: "dou1",
		meaning: "all",
	},
	{
		character: "然",
		pinyin: "ran2",
		meaning: "correct, so, like that",
	},
	{
		character: "沒",
		pinyin: "mei2",
		meaning: "haven't, there isn't, no",
	},
	{
		character: "日",
		pinyin: "ri4",
		meaning: "sun",
	},
	{
		character: "于",
		pinyin: "yu2",
		meaning: "in, at, for, to, from, by, than",
	},
	{
		character: "起",
		pinyin: "qi3",
		meaning: "rise, start",
	},
	{
		character: "還",
		pinyin: "hai2",
		meaning: "still, yet",
	},
	{
		character: "發",
		pinyin: "fa1",
		meaning: "deliver, develop, expand",
	},
	{
		character: "成",
		pinyin: "cheng2",
		meaning: "accomplish, succeed, to become",
	},
	{
		character: "事",
		pinyin: "shi4",
		meaning: "matter, thing, event, accident",
	},
	{
		character: "只",
		pinyin: "zhi3",
		meaning: "only, just",
	},
	{
		character: "作",
		pinyin: "zuo4",
		meaning: "do, make",
	},
	{
		character: "當",
		pinyin: "dang1",
		meaning: "serve as, act as",
	},
	{
		character: "想",
		pinyin: "xiang3",
		meaning: "think, feel, want",
	},
	{
		character: "看",
		pinyin: "kan4",
		meaning: "see, look at",
	},
	{
		character: "文",
		pinyin: "wen2",
		meaning: "language, literature",
	},
	{
		character: "無",
		pinyin: "wu2",
		meaning: "without, have not",
	},
	{
		character: "開",
		pinyin: "kai2",
		meaning: "open",
	},
	{
		character: "手",
		pinyin: "shou3",
		meaning: "hand",
	},
	{
		character: "十",
		pinyin: "shi2",
		meaning: "ten",
	},
	{
		character: "用",
		pinyin: "yong4",
		meaning: "use",
	},
	{
		character: "主",
		pinyin: "zhu3",
		meaning: "lord, master, live",
	},
	{
		character: "行",
		pinyin: "xing2",
		meaning: "go, OK",
	},
	{
		character: "方",
		pinyin: "fang2",
		meaning: "side",
	},
	{
		character: "又",
		pinyin: "you4",
		meaning: "again, both, and",
	},
	{
		character: "如",
		pinyin: "ru2",
		meaning: "like, as, as if",
	},
	{
		character: "前",
		pinyin: "qian2",
		meaning: "in front, previous, first, former",
	},
	{
		character: "所",
		pinyin: "suo3",
		meaning: "place",
	},
	{
		character: "本",
		pinyin: "ben3",
		meaning: "basis, origin, edition",
	},
	{
		character: "見",
		pinyin: "jian4",
		meaning: "appear to be, meet with",
	},
	{
		character: "經",
		pinyin: "jing1",
		meaning: "longitude",
	},
	{
		character: "頭",
		pinyin: "tou2",
		meaning: "head, top, first",
	},
	{
		character: "面",
		pinyin: "mian4",
		meaning: "face, surface",
	},
	{
		character: "公",
		pinyin: "gong1",
		meaning: "public, official",
	},
	{
		character: "同",
		pinyin: "tong2",
		meaning: "same",
	},
	{
		character: "三",
		pinyin: "san1",
		meaning: "three",
	},
	{
		character: "已",
		pinyin: "yi3",
		meaning: "stop, end",
	},
	{
		character: "老",
		pinyin: "lao3",
		meaning: "old",
	},
	{
		character: "從",
		pinyin: "cong2",
		meaning: "from, through",
	},
	{
		character: "動",
		pinyin: "dong4",
		meaning: "move",
	},
	{
		character: "兩",
		pinyin: "liang3",
		meaning: "two",
	},
	{
		character: "長",
		pinyin: "chang2",
		meaning: "long, length",
	},
}

