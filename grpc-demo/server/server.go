package main

import (
	"context"
	"flag"
	pb "github.com/zqddong/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

type GreeterServer struct {
}

// SayHello Unary RPC：一元 RPC
func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello.world"}, nil
}

// SayList Server-side streaming RPC：服务端流式 RPC
func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for n := 0; n <= 6; n++ {
		_ = stream.Send(&pb.HelloReply{Message: "hello.list " + string(n)})
	}
	return nil
}

// SayRecord Client-side streaming RPC：客户端流式 RPC
func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		rsp, err := stream.Recv()
		if err != io.EOF {
			return stream.SendAndClose(&pb.HelloReply{Message: "say.record"})
		}
		if err != nil {
			return err
		}
		log.Printf("Server SayRecord rsp: %v", rsp)
	}
	return nil
}

// SayRoute Bidirectional streaming RPC：双向流式 RPC
func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	n := 0
	for {
		_ = stream.Send(&pb.HelloReply{Message: "say.route"})

		rsp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++
		log.Printf("Server SayRoute rsp: %v", rsp)
	}
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}
