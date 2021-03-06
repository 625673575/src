package main

import (
	"net/rpc"
	"time"
	"fmt"
)

var client *rpc.Client
var err error
func callServer(i int) {

	var args = "hello rpc"
	var reply string
	err = client.Call("Echo.Hi", args, &reply)
	fmt.Print(reply)
	if err != nil {
		panic(err)
	}
}
func computeOnServer(i int){
	var reply float64
	err=client.Call("Echo.Add",[]float64{2.5,3.4},&reply)
	fmt.Print(reply)
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
		computeOnServer(i)
		time.Sleep(time.Second*1)
	}
}
