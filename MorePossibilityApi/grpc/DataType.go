package grpc

import "github.com/kaliwin/Needle/MorePossibilityApi/grpc/BurpMorePossibilityApi"

// GrpcHttpRequestData grpc http请求数据结构
type GrpcHttpRequestData struct {
	*BurpMorePossibilityApi.HttpReqData
}

// GrpcHttpResponseData grpc http响应数据结构
type GrpcHttpResponseData struct {
	*BurpMorePossibilityApi.HttpResData
}
