package parser

import (
	"log"
	"regexp"
	"stu/http/crawler/engine"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func CityListItemSaver(item engine.Item, saver chan interface{}) {
	cityInfo := item.Payload.(string)
	log.Printf("Got City:%s\n", cityInfo)
}

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		result.Items = append(result.Items, engine.Item{
			Payload:  string(m[2]),
			SaveFunc: CityListItemSaver,
		})
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}

	return result
}
