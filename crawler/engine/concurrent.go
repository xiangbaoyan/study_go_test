package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	//将request 加入schedule ，就是放到in中
	in := make(chan Request)

	//创造输出chan
	out := make(chan ParseResult)
	//创造in
	e.Scheduler.ConfigureMasterWorkChan(in)

	//异步执行 ，但是会卡死 10个
	//先建立起异步执行流程
	for i := 0; i < e.WorkerCount; i++ {
		//问题是结果放在 out 中怎么再处理
		createWorker(in, out)
	}
	//首次先加入一个主地址
	for k, r := range seeds {
		log.Printf("_join the scheduler:%d", k)
		//往in 中加入request
		e.Scheduler.Submit(r)
		log.Println("单次循环加入完毕")
	}
	log.Println("所有worker加入完毕")
	//收out 处理
	for {
		result := <-out
		log.Println("到达处理处...")

		itemCount := 0
		for _, item := range result.Items {
			log.Printf("Got item #%d:%v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}

	}
}

//是从request in中拿数据处理，返回到out （处理结果chan）中
//这个是一个过程 处理过程就是放在 goroutine 中执异步执行
//这是异步执行的过程，不是chan
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
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
