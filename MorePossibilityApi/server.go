package MorePossibilityApi

import (
	"context"
	"fmt"
	"github.com/kaliwin/Needle/MorePossibilityApi/grpc/BurpMorePossibilityApi"
	"github.com/kaliwin/Needle/PublicStandard/newsletter/StandardHttp"
	"google.golang.org/grpc"
	"log"
	"net"
)

// ApiTest api测试
type ApiTest struct {
	BurpMorePossibilityApi.UnimplementedHttpFlowHandlerServer
}

func (n ApiTest) HttpHandleRequestReceived(ctx context.Context, data *BurpMorePossibilityApi.HttpFlowReqData) (*BurpMorePossibilityApi.HttpRequestAction, error) {
	//TODO implement me
	return &BurpMorePossibilityApi.HttpRequestAction{Continue: true}, nil
}

func (n ApiTest) HttpHandleResponseReceived(ctx context.Context, data *BurpMorePossibilityApi.HttpFlowResData) (*BurpMorePossibilityApi.HttpResponseAction, error) {

	fmt.Println(data.GetHttpReqAndRes().GetReq().GetUrl())
	group, err := StandardHttp.BuildStandardHttpGroup(data.GetHttpReqAndRes(), nil)

	if err != nil {
		log.Println(err)

	} else {
		fmt.Println(group.Req.GetUrl())
		fmt.Println(group.Req.GetHead())
	}

	return &BurpMorePossibilityApi.HttpResponseAction{Continue: true}, nil
}

// NewGrpcServer 创建一个新的服务
func NewGrpcServer(address string, opt ...grpc.ServerOption) (GrpcServer, error) {
	server := GrpcServer{}
	server.serverStatus = ""
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return server, err
	}
	server.netList = listen
	server.Server = grpc.NewServer(opt...)
	return server, nil
}

// GrpcServer Grpc服务
type GrpcServer struct {
	Server       *grpc.Server
	netList      net.Listener
	serverStatus string
}

// Start 开启协程启动不会堵塞
func (s GrpcServer) Start() {
	go func() {
		err := s.Server.Serve(s.netList)
		if err != nil {
			s.serverStatus = err.Error()
			log.Println(err)
		}
	}()
}

// GetStatus 获取服务状态
func (s GrpcServer) GetStatus() error {
	if s.serverStatus != "" {
		return fmt.Errorf(s.serverStatus)
	}
	return nil
}
