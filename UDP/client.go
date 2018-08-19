// UDPClient project main.go
package main

import (
	"fmt"
	"net"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func main() {
	var buf [512]byte
	service := "47.75.241.11:5000"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	defer conn.Close()
	checkErr(err)
	rAddr := conn.RemoteAddr()
	n, err := conn.Write([]byte("Hello server!"))
	checkErr(err)
	n, err = conn.Read(buf[0:])
	checkErr(err)
	fmt.Println("Reply from server ", rAddr.String(), string(buf[0:n]))
	os.Exit(0)
}

