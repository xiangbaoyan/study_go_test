package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//创造输出chan
	out := make(chan ParseResult)
	e.Scheduler.Run()

	//异步执行 ，但是会卡死 10个
	//先建立起异步执行流程
	for i := 0; i < e.WorkerCount; i++ {
		//问题是结果放在 out 中怎么再处理
		//这里就建了10个
		createWorker(out, e.Scheduler)
	}
	//首次先加入一个主地址 ,把seeds 送进去
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	//收out 处理
	itemCount := 0
	for {
		result := <-out

		for _, item := range result.Items {
			log.Printf("Got item #%d:%v", itemCount, item)
			itemCount++

		}

		for _, request := range result.Requests {
			//在此处加入队列中
			e.Scheduler.Submit(request)
		}

	}

}

//是从request in中拿数据处理，返回到out （处理结果chan）中
//这个是一个过程 处理过程就是放在 goroutine 中执异步执行
//这是异步执行的过程，不是chan
func createWorker(out chan ParseResult, s Scheduler) {

	in := make(chan Request)

	//这是个gorouting ，执行完就释放了
	go func() {
		for {
			//从这对worker队列不断加入的
			//这个队列不断接受request
			//不断创建worker conn
			s.WorkerReady(in)
			//in 代表所有要处理的request
			request := <-in
			//在这进行处理
			result, err := worker(request)
			if err != nil {
				continue
			}
			//结果返回处理
			out <- result
		}
	}()
}
