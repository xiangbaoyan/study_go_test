package main

import (
	"fmt"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/config"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/rpcsupport"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawService{}))

}
