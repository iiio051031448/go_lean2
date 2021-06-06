package parser

import (
	"log"
	"stu/http/crawler/fetcher"
	"stu/http/crawler/model"
	"testing"
)

func TestParserProfile(t *testing.T) {
	contents, err := fetcher.Fetch("http://localhost:8080/mock/album.zhenai.com/u/1925392137109547791")
	//t.Logf("content: %s\n", contents)
	if err != nil {
		return
	}

	//t.Logf("content: %s\n", contents)

	resust := ParserProfile(contents, "心事痕迹酷酷猫")
	log.Printf("profile:%+v", resust.Items[0].(model.Profile))
}
