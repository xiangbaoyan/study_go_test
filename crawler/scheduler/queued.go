package scheduler

import (
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) ConfigureMasterWorkChan(chan engine.Request) {
	panic("implement me")
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {

		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {

			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
				//send r to worker
			//什么时候接受的，workReady
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
				//send next_request to w
				//这是一个worker conn 放进去request
			case activeWorker <- activeRequest:

				//分配出去处理了，但是怎么释放呢
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}

		}
	}()

}
