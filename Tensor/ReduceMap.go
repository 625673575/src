package tf
import (
	"net/rpc"
	"log"
	"net"
	"net/http"
	"fmt"
)
var client *rpc.Client
var err error
const (
	Address = "127.0.0.1:8086"
)
func  (t *Tensor) RemoteAdd(args *Tensor, reply *Tensor) error {
	//*reply =t.data[0]+args.data[0]
	t0:=Linear(2,5.5,1)
	t1:=Linear(3,6.5,1)
	err=client.Call("Echo.TensorAdd",[]Tensor{*t0,*t1},&reply)
	fmt.Print(reply)
	if err != nil {
		panic(err)
	}
	return nil
}
func Diag(){
	client, err = rpc.DialHTTP("tcp", "127.0.0.1:8086")
	if err != nil {
		println( err.Error())
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
func Start(protocal string, port string) {
	rpc.Register(new(Tensor))
	rpc.HandleHTTP()
	l, e := net.Listen(protocal, port)
	if e != nil {
		log.Panic("listen error: ", e)
	}

	http.Serve(l, nil)
}