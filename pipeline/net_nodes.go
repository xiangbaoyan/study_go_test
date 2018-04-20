package pipeline

import (
	"bufio"
	"fmt"
	"net"
)

func NetWorkSink(addr string, in <-chan int) {
	//这是在本地开server 的过程
	listener, e := net.Listen("tcp", addr)
	if e != nil {
		panic(e)
	}
	go func() {
		defer listener.Close()
		//开始接受链接，这是个等待的过程
		conn, e := listener.Accept()
		if e != nil {
			panic(e)
		}
		defer conn.Close()
		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		WriteSink(writer, in)

	}()

}

func NetWorkSource(addr string) <-chan int {
	out := make(chan int)
	go func() {
		conn, e := net.Dial("tcp", addr)
		if e != nil {
			panic(e)
		}

		fmt.Println("begin read net sorce:")
		r := ReadSource(bufio.NewReader(conn), -1)
		for v := range r {
			//fmt.Println("read out data:",v)
			out <- v
		}
		close(out)
	}()
	return out
}
