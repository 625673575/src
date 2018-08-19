package main

import(
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"ddz"
)
var ip string ="47.75.241.11:50001"//"0.0.0.0:50001"
func GrpcClient_Robot(){
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect :", err.Error())
	}else{
		fmt.Println("connect succ:"+conn.Target())
	}
	defer conn.Close()
	c :=ddz.NewRobotServiceClient(conn)
	lordCard := []byte{0x33,0x45,0x28,0x29,0x49,0x39}
	farmer1Card := []byte{0x23,0x13,0x43,0x46,0x26}
	farmer2Card := []byte{0x44,0x34,0x14}

	lastCards:=[]byte{0x1e,0x13,0x23,0x33,}
	r, err := c.Play(context.Background(), &ddz.RobotRequest{
		Playeridentity: 0,
		LordHandcard:lordCard,
		Farmer1Handcard:farmer1Card,
		Farmer2Handcard: farmer2Card,
		LastIdentity:2,
		LastPlaycard:lastCards,
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+x", r.Handcard)
	}
}
func GrpcClient_DealCard(){
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect :", err.Error())
	}else{
		fmt.Println("connect succ:"+conn.Target())
	}
	defer conn.Close()
	c :=ddz.NewDealCardServiceClient(conn)
	r, err := c.GetCard(context.Background(), &ddz.DealCardRequest{Params:[]byte{0,2}})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(r.Player0,r.Player1,r.Player2,r.Extra)
	}
}
func GrpcClient_Trustship(){
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect :", err.Error())
	}else{
		fmt.Println("connect succ:"+conn.Target())
	}
	defer conn.Close()
	c :=ddz.NewTrustshipServiceClient(conn)

	r, err := c.Ship(context.Background(), &ddz.TrustShipRequest{
		PlayerIdentity: 1,
		PlayerHandcard: []byte{0x44,0x34,0x14},
		LastIdentity: 2,
		LastPlaycard: []byte{0x1e,0x1e,0x13,0x23,0x33,},
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(r.Handcard)
	}
}
func main(){
	GrpcClient_DealCard()
	GrpcClient_Trustship()
	GrpcClient_Robot()
}
