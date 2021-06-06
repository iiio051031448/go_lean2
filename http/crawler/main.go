package main

import (
	"stu/http/crawler/engine"
	"stu/http/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList})
}
