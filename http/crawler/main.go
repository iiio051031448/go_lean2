package main

import (
	"stu/http/crawler/engine"
	"stu/http/crawler/scheduler"
	"stu/http/crawler/zhenai/parser"
)

func main() {
	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimpleScheduler{},
	//	WorkerCount: 100,
	//}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	//e.Run(engine.Request{
	//	Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList})
	//e.Run(engine.Request{
	//	Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun/gannan",
	//	ParseFunc: parser.ParseCity})
	e.Run(engine.Request{
		Url: "http://album.zhenai.com/u/1416899073",
		ParseFunc: func(c []byte) engine.ParseResult {
			return parser.ParserProfile(c, "偶偶偶")
		},
	})
}
