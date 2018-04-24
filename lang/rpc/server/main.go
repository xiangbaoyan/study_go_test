package main

import (
	"github.com/xiangbaoyan/study_go_test/lang/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//rpc 服务就起来了
func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accepted error: %+v", err)
		}
		go jsonrpc.ServeConn(conn)
	}

}
