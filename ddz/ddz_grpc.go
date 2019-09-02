package ddz

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strconv"
)

type ddzGrpcService struct {
	ip           string
	port         uint64
	robotService RobotServiceClient
	dealService  DealCardServiceClient
	conn         *grpc.ClientConn
}

func NewGrpcService(ip string, port uint64) (*ddzGrpcService,error) {
	r := new(ddzGrpcService)
	r.ip = ip
	r.port = port
	var err error
	r.conn, err = grpc.Dial(r.ip+":"+ strconv.Itoa(int(port)), grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect :", err.Error())
	}
	r.robotService = NewRobotServiceClient(r.conn)
	r.dealService = NewDealCardServiceClient(r.conn)
	return r,err
}

func (this *ddzGrpcService) Close() {
	this.conn.Close()
}

func (this *ddzGrpcService) Play(playerIdentity int32, lordCard []byte, f1Card []byte, f2Card []byte, lastIdentity int32, lastPlayCard []byte) ([]uint8, error) {
	r, err := this.robotService.Play(context.Background(), &RobotRequest{
		Playeridentity:  playerIdentity,
		LordHandcard:    lordCard,
		Farmer1Handcard: f1Card,
		Farmer2Handcard: f2Card,
		LastIdentity:    lastIdentity,
		LastPlaycard:    lastPlayCard,
	})
	if err != nil {
		return nil, err
	} else {
		return r.Handcard, err
	}
}

func (this *ddzGrpcService) Deal(param []byte) ([]uint8, []uint8, []uint8, []uint8, error) {
	if param == nil || len(param) == 0 {
		param = []byte{0, 0}
	}
	r, err := this.dealService.GetCard(context.Background(), &DealCardRequest{Params: param})
	if err != nil {
		return nil, nil, nil, nil, err
	} else {
		return r.Player0, r.Player1, r.Player2, r.Extra, err
	}
}
