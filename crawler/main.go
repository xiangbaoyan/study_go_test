package main

import (
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"github.com/xiangbaoyan/study_go_test/crawler/persist"
	"github.com/xiangbaoyan/study_go_test/crawler/scheduler"
	"github.com/xiangbaoyan/study_go_test/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
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
