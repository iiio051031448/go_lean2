package parser

import (
	"regexp"
	"strings"
	"stu/http/crawler/engine"
)

var cityRe = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var cityPageRe = regexp.MustCompile(`<span class="pager"><a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+/[^"]+)">([^>])+</a></span>`)

func ParseCity(contents []byte) engine.ParseResult {
	matchs := cityRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		name := string(m[2])
		url := string(m[1])
		urla := strings.Split(url, "/")
		if len(urla) < 3 {
			continue
		}
		id := urla[len(urla)-1]

		result.Items = append(result.Items, engine.Item{Payload: name})
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParserProfile(c, id, url, name)
			},
		})
	}

	matchs = cityPageRe.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		name := string(m[2])
		result.Items = append(result.Items, engine.Item{Payload: name})
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
		//log.Printf("P:%s\n", m[1])
	}

	return result
}
