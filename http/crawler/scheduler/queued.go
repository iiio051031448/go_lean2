package scheduler

import "stu/http/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
	go func() {
		var resquestQ []engine.Request
		var workQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(resquestQ) > 0 && len(workQ) > 0 {
				activeRequest = resquestQ[0]
				activeWorker = workQ[0]
			}

			select {
			case r := <-s.requestChan:
				resquestQ = append(resquestQ, r)
			case w := <-s.workerChan:
				workQ = append(workQ, w)
			case activeWorker <- activeRequest:
				resquestQ = resquestQ[1:]
				workQ = workQ[:1]
			}
		}
	}()
}
