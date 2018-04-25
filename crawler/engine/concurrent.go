package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(r Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	//我有一个worker 请问给我哪个channel
	WorkChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
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
		//开10个处理器不断执行
		e.createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	//首次先加入一个主地址 ,把seeds 送进去
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	//收out 处理
	for {
		result := <-out

		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
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
func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, read ReadyNotifier) {
	//这是个gorouting ，执行完就释放了
	go func() {
		for {
			//从这对worker队列不断加入的
			//这个队列不断接受request
			//不断创建worker conn
			//这里代表可以放同样的

			//问题是workCount 是什么含 义，大概是开10个处理器
			read.WorkerReady(in)
			//in 代表所有要处理的request
			request := <-in
			//在这进行处理
			result, err := e.RequestProcessor(request) //Worker// Worker(request)换成rpc
			if err != nil {
				continue
			}
			//结果返回处理
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicated(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
