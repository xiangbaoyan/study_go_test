package main

import (
	"github.com/xiangbaoyan/study_go_crawler/engine"
	"github.com/xiangbaoyan/study_go_crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
