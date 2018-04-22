package main

import (
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"github.com/xiangbaoyan/study_go_test/crawler/persist"
	"github.com/xiangbaoyan/study_go_test/crawler/scheduler"
	"github.com/xiangbaoyan/study_go_test/crawler/zhenai/parser"
)

func main() {

	e := engine.ConcurrentEngine{
		//在这传入的schedule 的类型
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
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
