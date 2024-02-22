package http

import (
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"net/http"
)

// ConvertHttp http转变接口 需要转到 httpReqData , 再由httpReqData转回自己的类型
// 转换为自己的时候要修改或构建自己的属性
type ConvertHttp interface {
	ConvertHttpReqDate() (*HttpStructureStandard.HttpReqData, error)
	ConvertHttpReqOwn(*HttpStructureStandard.HttpReqData) error

	ConvertHttpResDate() (*HttpStructureStandard.HttpResData, error)
	ConvertHttpResOwn(*HttpStructureStandard.HttpResData) error
}

// HttpClient http客户端
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// GrpcHttpReq Grpc下请求结构
type GrpcHttpReq interface {
	BurpHttpData
	GetUrl() string
	GetHttpVersion() string
	GetHttpReqService() *HttpStructureStandard.HttpReqService
}

// GrpcHttpRes Grpc下响应结构
type GrpcHttpRes interface {
	BurpHttpData
	GetHttpVersion() string
	//GetStatusCode() string
}

type GrpcHttpReqService interface {
	GetIp() string
	GetPort() int32
	GetSecure() bool
}

type BurpHttpData interface {
	GetData() []byte
	GetBodyIndex() int64
}

const (
	HTTP1         = "HTTP/1.1"
	ContentLength = "Content-Length"
)

// DefaultHeader 默认请求头
var DefaultHeader = http.Header{
	"User-Agent": []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.71 Safari/537.36"},
}

const (
	DefaultUA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.71 Safari/537.36"
)
