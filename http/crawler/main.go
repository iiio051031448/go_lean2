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
	e := engine.ConcurrentEngine2{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList})
}
