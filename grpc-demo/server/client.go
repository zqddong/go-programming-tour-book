package main

import (
	"context"
	"flag"
	pb "github.com/zqddong/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	_ = SayHello(client, &pb.HelloRequest{
		Name: "SayHello",
	})

	_ = SayList(client, &pb.HelloRequest{
		Name: "SayList",
	})

	_ = SayRecord(client, &pb.HelloRequest{
		Name: "SayRecord",
	})

	_ = SayRoute(client, &pb.HelloRequest{
		Name: "SayRoute",
	})
}

// SayHello Unary RPC：一元 RPC
func SayHello(client pb.GreeterClient, r *pb.HelloRequest) error {
	rsp, _ := client.SayHello(context.Background(), r)
	log.Printf("client.SayHello rsp: %s", rsp.Message)
	return nil
}

// SayList Server-side streaming RPC：服务端流式 RPC
func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("rsp: %v", resp)
	}
	return nil
}

// SayRecord Client-side streaming RPC：客户端流式 RPC
func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	for n := 0; n < 6; n++ {
		_ = stream.Send(r)
	}
	rsp, _ := stream.CloseAndRecv()
	log.Printf("Client SayRecord rsp: %v", rsp)
	return nil
}

// SayRoute Bidirectional streaming RPC：双向流式 RPC
func SayRoute(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRoute(context.Background())
	for n := 0; n <= 6; n++ {
		_ = stream.Send(r)
		rsp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("Client SayRoute rsp: %v", rsp)
	}

	_ = stream.CloseSend()

	return nil
}
