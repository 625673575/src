package ddz

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type DdzGrpcService struct {
	ip           string
	port         uint64
	robotService RobotServiceClient
	dealService  DealCardServiceClient
	conn         *grpc.ClientConn
	availability bool
}

func NewGrpcService(ip string, port uint64) (*DdzGrpcService, error) {
	r := new(DdzGrpcService)
	r.ip = ip
	r.port = port
	var err error
	r.conn, err = grpc.Dial(r.ip+":"+strconv.Itoa(int(port)), grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect :", err.Error())
	}
	r.robotService = NewRobotServiceClient(r.conn)
	r.dealService = NewDealCardServiceClient(r.conn)
	r.availability = true
	return r, err
}

func (this *DdzGrpcService) Close() {
	if this.conn != nil {
		this.conn.Close()
	}
}
func (this *DdzGrpcService) IsAvailable() bool {
	return this.availability
}
func (this *DdzGrpcService) String() string {
	return fmt.Sprintf("%s:%d", this.ip, this.port)
}
func (this *DdzGrpcService) GetIP() string {
	return this.ip
}
func (this *DdzGrpcService) GetPort() uint64 {
	return this.port
}

// playerIdentity 出牌者的身份 0是地主，1是地主下面的一个农民，2是再后面的农民
// lordCard , f1Card , f2Card 三个人的手牌
// lastIdentity 上一个出牌者的身份,如果是别人要不起自己出牌,则和playerIdentity一致
// lastPlayCard 上一个出牌者的牌
func (this *DdzGrpcService) Play(playerIdentity int32, lordCard []byte, f1Card []byte, f2Card []byte, lastIdentity int32, lastPlayCard []byte) ([]uint8, int, int, error) {
	r, err := this.robotService.Play(context.Background(), &RobotRequest{
		Playeridentity:  playerIdentity,
		LordHandcard:    lordCard,
		Farmer1Handcard: f1Card,
		Farmer2Handcard: f2Card,
		LastIdentity:    lastIdentity,
		LastPlaycard:    lastPlayCard,
	})
	if err != nil {
		errStatus := status.Convert(err)
		if errStatus.Code() == codes.Unavailable {
			this.availability = false
		}
		return nil, 0, 0, err
	} else {
		l := len(r.Handcard)
		if l == 0 {
			return r.Handcard, 0, 0, nil
		}
		return r.Handcard, int(r.Style.Style), int(r.Style.MaxVal), err
	}
}

func (this *DdzGrpcService) AnalyseCard(Cards [][]byte) *HandCardAnalysisReply {
	req := HandCardAnalysisRequest{
		Handcards: []*HandCardRequest{},
	}
	for _, v := range Cards {
		h := HandCardRequest{Cards: v}
		req.Handcards = append(req.Handcards, &h)
	}

	r, err := this.robotService.CardAnalysis(context.Background(), &req)
	if err != nil {
		return nil
	}
	return r
}

func (this *DdzGrpcService) Deal(param []byte) ([]uint8, []uint8, []uint8, []uint8, error) {
	if param == nil || len(param) == 0 {
		param = []byte{0, 0}
	}
	r, err := this.dealService.GetCard(context.Background(), &DealCardRequest{Params: param})
	if err != nil {
		errStatus := status.Convert(err)
		if errStatus.Code() == 14 {
			this.availability = false
		}
		return nil, nil, nil, nil, err
	} else {
		return r.Player0, r.Player1, r.Player2, r.Extra, err
	}
}
