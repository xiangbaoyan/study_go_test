package main

import (
	"fmt"
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"github.com/xiangbaoyan/study_go_test/crawler/modal"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/config"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/rpcsupport"
	"github.com/xiangbaoyan/study_go_test/lang/rpc"
	"testing"
	"time"
)

// type ItemSaverService has no exported methods of suitable type (hint: pass a pointer to value of that type)

//因为service 开在指针类型上所以报上面的错
func TestItemSaver(t *testing.T) {
	host := ":1234"
	//放一个server，出错后就不会注册和创建rpc 服务
	go serveRpc(host, "test1")
	time.Sleep(time.Second * 2)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/107157060",
		Type: "zhenai",
		Id:   "107157060",
		Payload: modal.Profile{
			Name:       "小红",
			Gender:     "女",
			Age:        0,
			Height:     160,
			Weight:     0,
			Income:     "3000元以下",
			Marriage:   "离异",
			Education:  "高中及以下",
			Occupation: "--",
			Hokou:      "--",
			Xinzuo:     "",
			House:      "--",
			Car:        "未购车",
		},
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil {
		panic(err)
	}
	//if err != nil || result != "ok"{
	//	t.Errorf("result: %s;err: %s",result,err)
	//}

}

func TestServer(t *testing.T) {
	go serveTest(":3456")
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(":3456")
	if err != nil {
		panic(err)
	}

	//conn, err := net.Dial("tcp", ":3456")
	//if err != nil {
	//	panic(err)
	//}
	//client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{4, 3}, &result)
	fmt.Println(result, err)

	//result :=""
	//err = client.Call("DemoService.Div", rpcdemo.Args{3, 4}, result)
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(result)
}
