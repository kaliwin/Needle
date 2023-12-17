package StandardHttp

import (
	"github.com/kaliwin/Needle/MorePossibilityApi/grpc/BurpMorePossibilityApi"
	"net/http"
)

// ConvertHttp http转变接口 需要转到 httpReqData , 再由httpReqData转回自己的类型
// 转换为自己的时候要修改或构建自己的属性
type ConvertHttp interface {
	ConvertHttpReqDate() (*BurpMorePossibilityApi.HttpReqData, error)
	ConvertHttpReqOwn(*BurpMorePossibilityApi.HttpReqData) error

	ConvertHttpResDate() (*BurpMorePossibilityApi.HttpResData, error)
	ConvertHttpResOwn(*BurpMorePossibilityApi.HttpResData) error
}

// HttpClient http客户端
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const (
	HTTP1         = "HTTP/1.1"
	ContentLength = "Content-Length"
)

// GrpcHttpReq Grpc下请求结构
type GrpcHttpReq interface {
	BurpHttpData
	GetUrl() string
	GetHttpVersion() string
	GetHttpReqService() *BurpMorePossibilityApi.HttpReqService
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

func TemporaryTransfer() {

}
