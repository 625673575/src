package main

import (
	"net"
	"log"
	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"
	"fmt"
	"RPC/remote"
)
import (
	"google.golang.org/grpc"
	"image/jpeg"
	"bytes"
	"github.com/vova616/screenshot"
	"time"
	"strings"
	"github.com/go-vgo/robotgo"
	cmd "RPC/command"
)

const port = ":50051"

type server struct{}

func (s *server) ExecCmd(ctx context.Context, in *remote.CmdRequest) (*remote.LogReply, error) {
	v,ok:=cmd.Commands[in.Cmd]
	log := new(remote.Log)
	if !ok {
		log.LogType=remote.LOG_TYPE_Error
		log.Content="No such command implementation"
		return &remote.LogReply{Logs: []*remote.Log{log}}, nil
	}
	log.LogType = remote.LOG_TYPE_Debug
	log.Content=	v(in.Args)
	return &remote.LogReply{Logs: []*remote.Log{log}}, nil
}
func (s *server) CaptureScreen(ctx context.Context, in *remote.CaptureParams) (*remote.ImageData, error) {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}
	o := new(jpeg.Options)
	o.Quality = (int)(in.Quality)
	emptyBuff := bytes.NewBuffer(nil)
	jpeg.Encode(emptyBuff, img, o) //img写入到buff
	return &remote.ImageData{Data: emptyBuff.Bytes(), TimeStamp: time.Now().Unix()}, nil
}

func (s *server) MouseClick(ctx context.Context, in *remote.MouseClickRequest) (*remote.Log, error) {
	robotgo.MoveMouse(int(in.GetX()), int(in.GetY()))
	switch in.GetButton() {
	case remote.MouseButtons_Left:
		robotgo.MouseClick("left", in.DoubleClick)
	case remote.MouseButtons_Middle:
		robotgo.MouseClick("mid", in.DoubleClick)
	case remote.MouseButtons_Right:
		robotgo.MouseClick("mid", in.DoubleClick)
	}
	log := new(remote.Log)
	log.LogType = remote.LOG_TYPE_Debug
	log.Content = "click position:" + in.String()
	fmt.Println(log.Content)
	return log, nil
}
func DelayTapKey( delay float32,keyCode []string){
	time.Sleep(time.Second*time.Duration(delay))
	p:=make([]interface{},len(keyCode))
	for i,v :=range keyCode{
		p[i]=v
	}
	robotgo.KeyTap(p...)
}
func (s *server) KeyTap(ctx context.Context, in *remote.KeyTapRequest) (*remote.Log, error){
	log := new(remote.Log)
	log.LogType = remote.LOG_TYPE_Debug
	log.Content =fmt.Sprintf("Press key %s after %.2f Seconds",strings.Join( in.KeyCode,"|") ,in.Delay)
	go DelayTapKey(in.Delay,in.KeyCode)
	return log,nil
}
func (s *server) KeyListTap(ctx context.Context, in *remote.KeyListTapRequest) (*remote.Log, error){	log := new(remote.Log)
	log.LogType = remote.LOG_TYPE_Debug
	log.Content =""
	for i,v :=range in.Keys{
		log.Content+= fmt.Sprintf("Press %dth key %s after %.2f Seconds",i,strings.Join( v.KeyCode,"|") ,v.Delay)
		go DelayTapKey(in.Keys[i].Delay,in.Keys[i].KeyCode)
	}
	return log,nil
}
func (s *server) GetMousePosition(ctx context.Context, in *remote.Nil) (*remote.MousePosition, error){
	x,y:=robotgo.GetMousePos()
	pos:=new(remote.MousePosition)
	pos.X=int32(x)
	pos.Y=int32(y)
	return pos,nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ser := new(server)
	remote.RegisterRemoteServiceServer(s, ser)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
