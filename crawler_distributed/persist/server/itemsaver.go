package main

import (
	"flag"
	"fmt"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/config"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/persist"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/rpcsupport"
	"github.com/xiangbaoyan/study_go_test/lang/rpc"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen")

func main() {

	flag.Parse()
	if *port == 0 {
		fmt.Println("must specified a port")
		return
	}
	//为了测试方便，可以把下边提取出方法
	//log.Fatal(serveRpc(":1234", "dating_profile"))

	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))

}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	//就因为这个地方 ！ 写成 = 出错 ，
	if err != nil {
		return err
	}
	//log.Printf("%+v",service)
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})

}
func serveTest(host string) error {
	return rpcsupport.ServeRpc(host, rpcdemo.DemoService{})

}
