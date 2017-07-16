package main

import (
	"net/rpc"
	"time"
)

var client *rpc.Client
var err error
func callServer(i int) {

	var args = "hello rpc"
	var reply string
	err = client.Call("Echo.Hi", args, &reply)

	if err != nil {
		panic(err)
	}
}
func diag(){
	client, err = rpc.DialHTTP("tcp", "127.0.0.1:8086")
	if err != nil {
		println( err.Error())
	}
}
func main(){
	diag()
	for i := 0; i < 5; i++ {
		callServer(i)
		time.Sleep(time.Second*1)
	}
}
