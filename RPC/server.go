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
)

const port = ":50001"

type server struct{}

func (s *server) ExecCmd(ctx context.Context, in *remote.CmdRequest) (*remote.LogReply, error) {
	log := new(remote.Log)
	log.LogType = remote.LOG_TYPE_Debug
	log.Content = "第一个log"
	fmt.Println(log)
	return &remote.LogReply{Logs: []*remote.Log{log}}, nil
}
func (s *server) CaptureScreen(ctx context.Context, in *remote.ScreenArea) (*remote.ImageData, error) {
	return &remote.ImageData{Data: []byte{12, 43, 234, 43}}, nil
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
