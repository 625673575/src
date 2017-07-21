package tf
import (
	"net/rpc"
	"log"
	"net"
	"net/http"
	"fmt"
)

const (
	Address = "127.0.0.1:8086"
)
func  (t *Tensor) RemoteAdd(args *Tensor, reply *float64) error {
	*reply =t.data[0]+args.data[0]
	fmt.Println(*reply)
	return nil
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