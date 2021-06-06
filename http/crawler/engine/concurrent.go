package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}

		for _, nr := range result.Requests {
			e.Scheduler.Submit(nr)
		}
	}
}

var workIdNext int = 0

func createWorker(in chan Request, out chan ParseResult) {
	workId := workIdNext2
	workIdNext++
	go func() {
		for {
			r := <-in
			parseResult, err := worker(r, workId)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()

}
