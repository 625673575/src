package main

import (
	"net/rpc"
	"log"
	"net"
	"net/http"
	"time"
	"fmt"
)

const (
	Address = "127.0.0.1:8086"
)

type Echo int
func (t *Echo) Hi(args string, reply *string) error {
	*reply = "echo:" + args
	fmt.Println(*reply)
	return nil
}
func Start(protocal string, port string) {
	rpc.Register(new(Echo))
	rpc.HandleHTTP()
	l, e := net.Listen(protocal, port)
	if e != nil {
		log.Panic("listen error: ", e)
	}
	http.Serve(l, nil)
}
func main() {
	go Start("tcp", ":8086")
	//time.Sleep(time.Second * 5)
	//for i := 0; i < 5; i++ {
	//	call(i)
	//	time.Sleep(time.Second*1)
	//}
	time.Sleep(time.Hour)
}
func call(i int) {
	client, err := rpc.DialHTTP("tcp", Address)
	if err != nil {
		println(i, err.Error())
	}
	var args = "hello rpc"
	var reply string
	err = client.Call("Echo.Hi", args, &reply)

	if err != nil {
		panic(err)
	}
}
