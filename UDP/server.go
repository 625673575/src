// UDPServer project main.go
package main

import (
	"fmt"
	"net"
)

func main() {
	service := ":5000"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkErr(err)
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, rAddr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			return
		}
		fmt.Println("Receive from client", rAddr.String(), string(buf[0:n]))
		_, err2 := conn.WriteToUDP([]byte("Welcome client!"), rAddr)
		if err2 != nil {
			return
		}
	}
}
