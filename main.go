package main

import (
	"context"
	"fmt"
	"github.com/kaliwin/Needle/MorePossibilityApi"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"google.golang.org/grpc"
)

type Tc struct {
	ClientId string `json:"clientId"`
	Data     string `json:"data"`
	Status   int    `json:"status"`
}

func (t Tc) HttpFlowOut(c context.Context, reqAndRes *HttpStructureStandard.HttpReqAndRes) (*HttpStructureStandard.Str, error) {

	fmt.Println(reqAndRes.GetReq().GetUrl())

	return &HttpStructureStandard.Str{Name: ""}, nil
}

func main() {

	server, err := MorePossibilityApi.NewBurpGrpcServer(":9001", grpc.MaxRecvMsgSize(200*1024*1024))
	if err != nil {
		panic(err)
	}

	server.RegisterHttpFlowOut(&Tc{})

	server.Start()

}
