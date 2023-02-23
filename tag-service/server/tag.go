package server

import (
	"context"
	"encoding/json"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/zqddong/go-programming-tour-book/tag-service/internal/middleware"
	"github.com/zqddong/go-programming-tour-book/tag-service/pkg/bapi"
	"github.com/zqddong/go-programming-tour-book/tag-service/pkg/errcode"
	pb "github.com/zqddong/go-programming-tour-book/tag-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type TagServer struct {
	auth *Auth
}

type Auth struct{}

func (a *Auth) GetAppKey() string {
	return "go-programming-tour-book"
}

func (a *Auth) GetAppSecret() string {
	return "eddycjy"
}

func (a *Auth) Check(ctx context.Context) error {
	md, _ := metadata.FromIncomingContext(ctx)

	var appKey, appSecret string
	if value, ok := md["app_key"]; ok {
		appKey = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}
	if appKey != a.GetAppKey() || appSecret != a.GetAppSecret() {
		return errcode.TogRPCError(errcode.Unauthorized)
	}

	return nil
}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	//if err := t.auth.Check(ctx); err != nil {
	//	return nil, err
	//}

	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}

	return &tagList, nil
}

// gRPC 服务内调
func (t *TagServer) GetTagList2(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	clientConn, err := GetClientConn(ctx, "localhost:8001",
		[]grpc.DialOption{grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				middleware.UnaryContextTimeout(),
				middleware.ClientTracing(),
			),
		)},
	)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	defer clientConn.Close()
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	rsp, err := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"})
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return rsp, nil
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
