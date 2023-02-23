#### Protobuf 

1. protoc 安装
```shell
$ wget https://github.com/google/protobuf/releases/download/v3.11.2/protobuf-all-3.11.2.zip
$ unzip protobuf-all-3.11.2.zip && cd protobuf-3.11.2/
$ ./configure
$ make
$ make install

protoc --version # 执行报错 执行命令 ldconfig
```

2. protoc 插件安装
```shell
$ GIT_TAG="v1.3.2"
$ go get -d -u github.com/golang/protobuf/protoc-gen-go
$ git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
$ go install github.com/golang/protobuf/protoc-gen-go

# 二进制文件 protoc-gen-go 移动到 bin 目录下，让其可以直接运行 protoc-gen-go 执行
$ mv $GOPATH/bin/protoc-gen-go /usr/local/go/bin/
```

3. 生成 proto 文件
```shell
$ protoc --go_out=plugins=grpc:. ./proto/*.proto
```

#### gRPC 

1. 包
```shell
$ go get -u google.golang.org/grpc@v1.29.1
```

2. 调试 gRPC 接口
```shell
$ go get github.com/fullstorydev/grpcurl
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl

# 安装后
$ grpcurl -plaintext localhost:8001 list
grpc.reflection.v1alpha.ServerReflection
proto.TagService

$ grpcurl -plaintext localhost:8001 list proto.TagService
proto.TagService.GetTagList

$ grpcurl -plaintext -d '{"name":"Go"}' localhost:8001 proto.TagService.GetTagList
```


- gRPC 一共支持四种调用方式，分别是：
    1. Unary RPC：一元 RPC
    2. Server-side streaming RPC：服务端流式 RPC。
    3. Client-side streaming RPC：客户端流式 RPC。
    4. Bidirectional streaming RPC：双向流式 RPC。

- gRPC 在建立连接之前，客户端/服务端都会发送连接前言（Magic+SETTINGS），确立协议和配置项。

- gRPC 在传输数据时，是会涉及滑动窗口（WINDOW_UPDATE）等流控策略的。

- 传播 gRPC 附加信息时，是基于 HEADERS 帧进行传播和设置；而具体的请求/响应数据是存储的 DATA 帧中的。

- gRPC 请求/响应结果会分为 HTTP 和 gRPC 状态响应（grpc-status、grpc-message）两种类型。

- 客户端发起 PING，服务端就会回应 PONG，反之亦可。