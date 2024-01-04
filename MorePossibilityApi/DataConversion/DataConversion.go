package DataConversion

import "github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard"

// GrpcHttpReqDataToStandardHttpRequest 将grpc的http请求数据转换为标准的http请求数据
func GrpcHttpReqDataToStandardHttpRequest() (HttpStructureStandard.HttpRequest, error) {

	return HttpStructureStandard.HttpRequest{}, nil
}

// GrpcHttpResDataToStandardHttpResponse 将grpc的http响应数据转换为标准的http响应数据
func GrpcHttpResDataToStandardHttpResponse() (HttpStructureStandard.HttpResponse, error) {

	return HttpStructureStandard.HttpResponse{}, nil
}
