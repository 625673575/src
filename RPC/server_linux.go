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
	"time"
	"strings"
	"RPC/file"
	"io/ioutil"
	"os"
)

const port = ":50051"
const MAX_TEXT_LEN=2097152

type server struct{}

func (s *server) ExecCmd(ctx context.Context, in *remote.CmdRequest) (*remote.LogReply, error) {
	return &remote.LogReply{Logs: []*remote.Log{}}, nil
}
func (s *server) CaptureScreen(ctx context.Context, in *remote.CaptureParams) (*remote.ImageData, error) {
	return &remote.ImageData{Data:[]byte{}, TimeStamp: time.Now().Unix()}, nil
}

func (s *server) MouseClick(ctx context.Context, in *remote.MouseClickRequest) (*remote.Log, error) {
	log := new(remote.Log)
	log.LogType = remote.LOG_TYPE_Debug
	log.Content = "not implement in linux"
	fmt.Println(log.Content)
	return log, nil
}
func (s *server) KeyTap(ctx context.Context, in *remote.KeyTapRequest) (*remote.Log, error) {
	log := new(remote.Log)
	log.LogType = remote.LOG_TYPE_Debug
	log.Content = "not implement in linux"
	fmt.Println(log.Content)
	return log, nil
}
func (s *server) KeyListTap(ctx context.Context, in *remote.KeyListTapRequest) (*remote.Log, error) {
	log := new(remote.Log)
	log.LogType = remote.LOG_TYPE_Debug
	log.Content = "not implement in linux"
	fmt.Println(log.Content)
	return log, nil
}
func (s *server) GetMousePosition(ctx context.Context, in *remote.Nil) (*remote.MousePosition, error) {
p:=new(remote.MousePosition)
	return p, nil
}

func (s *server) ListDir(ctx context.Context, in *remote.CmdRequest) (*remote.FileEntrys, error) {
	r,err:=file.ListDirAndFile(in.Cmd)
	if err!=nil{
		r.Count=0
		r.Entrys=nil
	}
	return &r,nil
}
func (s *server) ListDisk(ctx context.Context, in *remote.Nil) (*remote.Log, error) {
	disks := file.ListDisk()
	log := new(remote.Log)
	log.LogType = remote.LOG_TYPE_Debug
	log.Content = strings.Join(disks, " ")
	return log, nil
}
func (s *server) ReadFile(ctx context.Context, in *remote.CmdRequest) (*remote.FileData, error) {
	r:=new(remote.FileData)
	var err error
	r.FullName=in.Cmd
	r.Data,err=file.ReadBytes(r.FullName)
	if err!=nil{
		r.FullName="读取文件失败"
	}
	return r, nil
}
func (s *server) ReadText(ctx context.Context, in *remote.CmdRequest) (*remote.Log, error)      {
	var err error
	log := new(remote.Log)
	log.LogType = remote.LOG_TYPE_Debug
	if file.GetFileSize(in.Cmd)>MAX_TEXT_LEN{
		log.LogType=remote.LOG_TYPE_Error
		log.Content=fmt.Sprintf("%s文件体积大于%dkb",in.Cmd,MAX_TEXT_LEN/1024)
		return log,nil
	}
	log.Content,err =file.ReadString(in.Cmd)
	if err!=nil{
		log.LogType=remote.LOG_TYPE_Error
		log.Content=fmt.Sprintf("读取文本%s失败",in.Cmd)
	}
	return log, nil
}
func (s *server) CreateFile(ctx context.Context, in *remote.FileData) (*remote.Log, error)      {
	log := new(remote.Log)
	log.LogType = remote.LOG_TYPE_Debug
	err:= ioutil.WriteFile(in.FullName,in.Data,os.ModePerm)
	if err==nil{
		log.Content=fmt.Sprintf("创建文件%s成功",in.FullName)
	}
	return log, nil
}
func (s *server) CreateText(ctx context.Context, in *remote.CmdRequest) (*remote.Log, error)    {
	log := new(remote.Log)
	if len(in.Args)<1{
		log.LogType = remote.LOG_TYPE_Error
		log.Content="参数不正确,arg[0]=文本内容"
	}
	err:= ioutil.WriteFile(in.Cmd,[]byte(in.Args[0]),os.ModePerm)
	if err==nil{
		log.LogType= remote.LOG_TYPE_Debug
		log.Content=fmt.Sprintf("创建文本%s成功",in.Cmd)
	}
	return log, nil
}
func (s *server) Rename(ctx context.Context, in *remote.CmdRequest) (*remote.Log, error)        {
	log := new(remote.Log)
	if len(in.Args)<1{
		log.LogType = remote.LOG_TYPE_Error
		log.Content="参数不正确,cmd=文件名,arg[0]=新文件名称"
	}
	succ:=file.Rename(in.Cmd,in.Args[0])
	if succ{
		log.LogType= remote.LOG_TYPE_Debug
		log.Content=fmt.Sprintf("修改文件%s文件名为%s",in.Cmd,in.Args[0])
	}
	return log, nil
}
func (s *server) DeleteFile(ctx context.Context, in *remote.CmdRequest) (*remote.Log, error)    {
	log := new(remote.Log)
	succ:=os.Remove(in.Cmd)
	if succ==nil{
		log.LogType= remote.LOG_TYPE_Debug
		log.Content=fmt.Sprintf("删除文件%s成功",in.Cmd)
	}else{
		log.LogType= remote.LOG_TYPE_Error
		log.Content=fmt.Sprintf("删除文件%s失败",in.Cmd)
	}
	return log, nil
}
func (s *server) DeleteFolder(ctx context.Context, in *remote.CmdRequest) (*remote.Log, error)  {
	log := new(remote.Log)
	succ:=os.Remove(in.Cmd)
	if succ==nil{
		log.LogType= remote.LOG_TYPE_Debug
		log.Content=fmt.Sprintf("删除文件夹%s成功",in.Cmd)
	}else{
		log.LogType= remote.LOG_TYPE_Error
		log.Content=fmt.Sprintf("删除文件夹%s失败",in.Cmd)
	}
	return log, nil
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
