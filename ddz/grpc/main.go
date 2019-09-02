package main

import (
	"ddz"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"sync"
)

var ip string = "47.75.241.11:50001"

func TypeCall(){
	service,err := ddz.NewGrpcService("47.75.241.11", 50001)
	if err!=nil{
		panic(err)
	}
	l, f1, f2, extra, err := service.Deal(nil)
	if err == nil {
		fmt.Printf("Lord = %v,Farmer1 = %v,Farmer2 = %v,Extra = %v \n", l, f1, f2, extra)
	}
	lordCard := []byte{0x33, 0x25, 0x45, 0x28, 0x29, 0x49, 0x39}
	farmer1Card := []byte{0x23, 0x13, 0x43, 0x46, 0x26}
	farmer2Card := []byte{0x44, 0x34, 0x14}
	lastCard := []byte{0x1e, 0x1e, 0x13, 0x23, 0x33}
	r, err := service.Play(0,lordCard,farmer1Card,farmer2Card,2,lastCard)
	if err==nil{
		fmt.Printf("AI出牌 %v \n",r)
	}
}


func GrpcClient_Robot() {
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect :", err.Error())
	}
	defer conn.Close()
	c := ddz.NewRobotServiceClient(conn)

	lordCard := []byte{0x33, 0x25, 0x45, 0x28, 0x29, 0x49, 0x39}
	farmer1Card := []byte{0x23, 0x13, 0x43, 0x46, 0x26}
	farmer2Card := []byte{0x44, 0x34, 0x14}

	lastCards := []byte{0x1e, 0x1e, 0x13, 0x23, 0x33}
	r, err := c.Play(context.Background(), &ddz.RobotRequest{
		Playeridentity:  0,
		LordHandcard:    lordCard,
		Farmer1Handcard: farmer1Card,
		Farmer2Handcard: farmer2Card,
		LastIdentity:    2,
		LastPlaycard:    lastCards,
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(r.Handcard)
	}
}

var connDealCardService ddz.DealCardServiceClient

func GrpcClient_DealCard() {
	r, err := connDealCardService.GetCard(context.Background(), &ddz.DealCardRequest{Params: []byte{0, 2}})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(r.Player0, r.Player1, r.Player2, r.Extra)
	}
}
func GrpcClient_Trustship() {
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect :", err.Error())
	}
	defer conn.Close()
	c := ddz.NewTrustshipServiceClient(conn)

	r, err := c.Ship(context.Background(), &ddz.TrustShipRequest{
		PlayerIdentity: 1,
		PlayerHandcard: []byte{0x44, 0x34, 0x14},
		LastIdentity:   1,
		LastPlaycard:   []byte{0x1e, 0x1e, 0x13, 0x23, 0x33},
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(r.Handcard)
	}
}

var wg sync.WaitGroup

func main() {
	GrpcClient_Trustship()
	GrpcClient_Robot()
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect :", err.Error())
	}
	defer conn.Close()
	connDealCardService = ddz.NewDealCardServiceClient(conn)
	i := 0
	for ; ;
	{
		i++
		log.Println(i)
		GrpcClient_DealCard()
	}
}
