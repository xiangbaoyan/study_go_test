package main

import (
	"fmt"
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"github.com/xiangbaoyan/study_go_test/crawler/scheduler"
	"github.com/xiangbaoyan/study_go_test/crawler/zhenai/parser"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/config"
	itemsaver "github.com/xiangbaoyan/study_go_test/crawler_distributed/persist/client"
	worker "github.com/xiangbaoyan/study_go_test/crawler_distributed/worker/client"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		//在这传入的schedule 的类型
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		//这个类型是返回值 chan
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	//request := engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//}

	request := engine.Request{
		Url:    "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	}
	e.Run(request)
}
