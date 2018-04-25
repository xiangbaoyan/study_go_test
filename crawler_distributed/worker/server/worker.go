package main

import (
	"flag"
	"fmt"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/rpcsupport"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/worker"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specified port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawService{}))

}
