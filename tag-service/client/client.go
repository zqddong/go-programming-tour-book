package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/naming"
	pb "github.com/zqddong/go-programming-tour-book/tag-service/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

type Auth struct {
	AppKey    string
	AppSecret string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_key": a.AppKey, "app_secret": a.AppSecret}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return false
}
func main() {
	//auth := Auth{
	//	AppKey:    "go-programming-tour-boo2k",
	//	AppSecret: "eddycjy",
	//}

	ctx := context.Background()
	//opts := []grpc.DialOption{grpc.WithPerRPCCredentials(&auth)}
	//clientConn, _ := GetClient(ctx, "localhost:8004", opts)
	//newCtx := metadata.AppendToOutgoingContext(ctx, "eddycjy", "Go 语言编程之旅")
	//clientConn, err := GetClient(newCtx, "localhost:8001", []grpc.DialOption{
	//	grpc.WithUnaryInterceptor(
	//		grpc_middleware.ChainUnaryClient(
	//			middleware.UnaryContextTimeout(),
	//			middleware.ClientTracing(),
	//		),
	//	),
	//	grpc.WithStreamInterceptor(
	//		grpc_middleware.ChainStreamClient(
	//			middleware.StreamContextTimeout(),
	//		),
	//	),
	//	// gRPC 自定义认证
	//	grpc.WithPerRPCCredentials(&auth),
	//})

	clientConn, err := GetClient2(ctx, "tag-service", nil)

	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer clientConn.Close()

	tagServiceClient := pb.NewTagServiceClient(clientConn)
	rsp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"})

	log.Printf("rps:%v", rsp)
}

func GetClient(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}

func GetClient2(ctx context.Context, serviceName string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	config := clientv3.Config{
		DialTimeout: time.Second * 60,
		Endpoints:   []string{"http://localhost:2379"},
	}

	cli, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}

	r := &naming.GRPCResolver{Client: cli}
	target := fmt.Sprintf("/etcdv3://go-programing-tour/grpc/%s", serviceName)

	opts = append(opts, grpc.WithInsecure(), grpc.WithBalancer(grpc.RoundRobin(r)), grpc.WithBlock())
	return grpc.DialContext(ctx, target, opts...)
}
