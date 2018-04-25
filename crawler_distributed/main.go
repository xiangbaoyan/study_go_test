package main

import (
	"flag"
	"fmt"
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"github.com/xiangbaoyan/study_go_test/crawler/scheduler"
	"github.com/xiangbaoyan/study_go_test/crawler/zhenai/parser"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/config"
	itemsaver "github.com/xiangbaoyan/study_go_test/crawler_distributed/persist/client"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/rpcsupport"
	worker "github.com/xiangbaoyan/study_go_test/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String(
		"worker_hosts", "", "worker hosts (commma seperated)")
)

func main() {

	flag.Parse()

	itemChan, err := itemsaver.ItemSaver(
		fmt.Sprintf(":%d", *itemSaverHost))
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ", "))

	processor := worker.CreateProcessor(pool)
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

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connnected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}

	}

	//分发
	out := make(chan *rpc.Client)
	go func() {
		//要让其不断分发
		for {
			//轮流分发
			for _, client := range clients {
				out <- client
			}
		}

	}()

	return out

}
