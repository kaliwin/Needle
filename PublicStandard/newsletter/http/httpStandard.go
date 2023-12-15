package http

import (
	"github.com/kaliwin/Needle/PublicStandard/newsletter/http/grpc/GrpcHttpStandard"
	"net/http"
)

// ConvertHttp http转变接口 需要转到 httpReqData , 再由httpReqData转回自己的类型
// 转换为自己的时候要修改或构建自己的属性
type ConvertHttp interface {
	ConvertHttpReqDate() (*GrpcHttpStandard.HttpReqData, error)
	ConvertHttpReqOwn(*GrpcHttpStandard.HttpReqData) error

	ConvertHttpResDate() (*GrpcHttpStandard.HttpResData, error)
	ConvertHttpResOwn(*GrpcHttpStandard.HttpResData) error
}

// ClientHttp http客户端
type ClientHttp interface {
	Do(req *http.Request) (*http.Response, error)
}

const (
	HTTP1         = "HTTP/1.1"
	ContentLength = "Content-Length"
)
