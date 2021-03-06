package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	log.Printf("listen on port: %s", host)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accepted error: %+v", err)
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClient(host string) (*rpc.Client, error) {

	//conn, err := net.Dial("tcp", ":3456")
	//if err != nil {
	//	panic(err)
	//}
	//client := jsonrpc.NewClient(conn)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil

}
