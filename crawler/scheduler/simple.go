//安排处理器
package scheduler

import (
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"log"
)

type SimpleScheduler struct {
	//请求处理worker chan 但不是worker 是request chan
	//和in的关系 其实就是in(外部传来的)
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkChan(c chan engine.Request) {
	s.workerChan = c
}

//卡死的地方需要用goroutine 解决
func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		log.Println("_enter into submit")
		s.workerChan <- request
		log.Println("加入完成")

	}()

}
