package main

import (
	"stu/http/crawler/engine"
	"stu/http/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "https://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList})
}
