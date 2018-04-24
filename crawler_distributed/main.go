package main

import (
	"fmt"
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"github.com/xiangbaoyan/study_go_test/crawler/scheduler"
	"github.com/xiangbaoyan/study_go_test/crawler/zhenai/parser"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/config"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := client.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		//在这传入的schedule 的类型
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		//这个类型是返回值 chan
		ItemChan: itemChan,
	}
	//request := engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//}
	request := engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/shanghai",
		ParseFunc: parser.ParseCity,
	}
	e.Run(request)
}
