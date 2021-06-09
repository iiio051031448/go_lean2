package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkReady(chan Request)
}

var visitedUrls = make(map[string]bool)

func isDuplicate(r Request) bool {
	if visitedUrls[r.Url] {
		return true
	}

	visitedUrls[r.Url] = true
	return false
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
			//log.Printf("Got item #%d: %v", itemCount, item)
			if item.HandleFunc != nil {
				item.HandleFunc(item)
			}

			itemCount++
		}

		for _, nr := range result.Requests {
			if isDuplicate(nr) {
				log.Printf("Duplicate request: %s\n", nr.Url)
				continue
			}
			e.Scheduler.Submit(nr)
		}
	}
}

var workIdNext2 int = 0

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	workIdNext2++
	go func(wId int) {
		for {
			log.Printf("Work [%d] is Ready\n", wId)
			ready.WorkReady(in)
			r := <-in
			parseResult, err := worker(r, wId)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}(workIdNext2)

}
