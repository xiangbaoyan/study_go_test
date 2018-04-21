package main

import (
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"github.com/xiangbaoyan/study_go_test/crawler/scheduler"
	"github.com/xiangbaoyan/study_go_test/crawler/zhenai/parser"
)

func main() {

	e := engine.ConcurrentEngine{
		//在这传入的schedule 的类型
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	request := engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	}
	e.Run(request)
}
