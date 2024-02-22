package grpc

import (
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
)

// GrpcHttpRequestData grpc http请求数据结构
type GrpcHttpRequestData struct {
	*HttpStructureStandard.HttpReqData
}

// GrpcHttpResponseData grpc http响应数据结构
type GrpcHttpResponseData struct {
	*HttpStructureStandard.HttpResData
}
