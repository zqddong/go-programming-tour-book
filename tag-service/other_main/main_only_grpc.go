package main

import (
	"flag"
	pb "github.com/zqddong/go-programming-tour-book/tag-service/proto"
	"github.com/zqddong/go-programming-tour-book/tag-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8001", "gRPC 启动端口号")
	flag.Parse()
}

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("s.Server err: %v", err)
	}

}
