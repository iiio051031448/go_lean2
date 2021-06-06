package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	WorkReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
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

var workIdNext2 int = 0

func createWorker(in chan Request, out chan ParseResult, s Scheduler) {
	workIdNext2++
	go func(wId int) {
		for {
			log.Printf("Work [%d] is Ready\n", wId)
			s.WorkReady(in)
			r := <-in
			parseResult, err := worker(r, wId)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}(workIdNext2)

}
