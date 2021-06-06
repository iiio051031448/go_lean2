package parser

import (
	"regexp"
	"stu/http/crawler/engine"
)

const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParserProfile(c, name)
			},
		})
	}

	return result
}
