//安排处理器
package scheduler

import (
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
)

type SimpleScheduler struct {
	//请求处理worker chan 但不是worker 是request chan
	//和in的关系 其实就是in(外部传来的)
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

//卡死的地方需要用goroutine 解决
//卡死原因，如果in满了后，就停下了，阻止了下一次循环，就不能把request再加入
//每个request 建立一个gorouting
func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		//放入in 如果in满（因为就建了10个），放不进去，没有gorouting 就卡死了
		s.workerChan <- request
	}()

}
