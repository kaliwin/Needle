package http

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/OriginalMessage"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	http2 "github.com/kaliwin/Needle/network/http"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Request struct {
	RequestOriginalMessage *OriginalMessage.RequestOriginalMessage // 原始请求消息
	scheme                 string                                  // 协议
	url                    *url.URL
}

// GetScheme 获取协议
func (r *Request) GetScheme() string {
	return r.scheme
}

// SetScheme 设置协议
func (r *Request) SetScheme(Scheme string) {
	r.scheme = Scheme
}

// GetUrl 获取url
func (r *Request) GetUrl() (string, error) {
	if r.url == nil {
		return "", errors.New("url is nil")
	}
	return r.url.String(), nil
}

// SetUrl 设置url
func (r *Request) SetUrl(url *url.URL) {
	r.scheme = url.Scheme
	r.url = url
}

// GetMethod 获取请求方法
func (r *Request) GetMethod() string {
	return r.RequestOriginalMessage.RequestLine.Method
}

// GetPath 获取请求路径 包含?后的参数
func (r *Request) GetPath() string {
	return r.RequestOriginalMessage.RequestLine.Path
}

// GetHead 获取头部实例
func (r *Request) GetHead() http.Header {
	return http.Header(r.RequestOriginalMessage.Head)
}

// GetHeadValue 获取头部值
func (r *Request) GetHeadValue(key string) []string {
	return r.RequestOriginalMessage.Head[key]
}

// GetBody 获取请求体
func (r *Request) GetBody() []byte {
	return r.RequestOriginalMessage.Body
}

// SetMethod 设置请求方法
func (r *Request) SetMethod(method string) {
	r.RequestOriginalMessage.RequestLine.Method = method
}

// SetPath 设置请求路径 包含?后的参数
func (r *Request) SetPath(path string) {
	r.RequestOriginalMessage.RequestLine.Path = path
}

// SetNewHead 设置新的头部
func (r *Request) SetNewHead(head map[string][]string) {
	r.RequestOriginalMessage.Head = head
}

// SetHead 会覆盖原有的头键值
func (r *Request) SetHead(key string, value []string) {
	r.RequestOriginalMessage.Head[key] = value
}

// AddHead 增加头部键值
func (r *Request) AddHead(key string, value string) {
	r.RequestOriginalMessage.Head[key] = append(r.RequestOriginalMessage.Head[key], value)
}

// SetBody 设置请求体
func (r *Request) SetBody(body []byte) {
	r.RequestOriginalMessage.Body = body
}

// BuildMessage 构建请求报文
func (r *Request) BuildMessage() []byte {
	return r.RequestOriginalMessage.BuildMessage()
}

// BuildRawHttpRequest 构建原始请求 没有处理URL
// 强制使用HTTP/1.1
func (r *Request) BuildRawHttpRequest() (*http.Request, error) {
	r.RequestOriginalMessage.RequestLine.HttpVersion = http2.HTTP1

	request, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(r.BuildMessage())))
	if err != nil {
		return nil, err
	}
	return request, err
}

// BuildHttpRequest 构建可发送的请求 没有设置url 会从原始请求中获取到Host确定目标 但是需要指定协议
// url 不为空 会直接使用url
func (r *Request) BuildHttpRequest(scheme string) (*http.Request, error) {
	request, err := r.BuildRawHttpRequest()
	if err != nil {
		return request, err
	}

	if r.url != nil {
		request.URL = r.url
		return request, nil
	}
	if scheme != "" {
		r.scheme = scheme
	}

	buildUrl, err := OriginalMessage.BuildUrl(r.scheme, request.Host, request.RequestURI)
	if err != nil {
		return request, err
	}
	request.URL = buildUrl
	request.RequestURI = ""
	return request, nil

}

// BuildGrpcRequest 构建grpc请求实例
func (r *Request) BuildGrpcRequest() *HttpStructureStandard.HttpReqData {

	message := r.BuildMessage()
	i := len(r.GetBody())

	ip := ""
	port := 0

	if r.url != nil {
		if r.url.Host != "" {
			str := strings.Split(r.url.Host, ":")
			if str[0] != "" {
				ip = str[0]
			}
			if len(str) > 1 {
				port, _ = strconv.Atoi(str[1])
			}
		}
		if r.scheme == "" {
			r.scheme = r.url.Scheme
		}

	}
	if ip == "" {
		host := r.GetHead().Get("Host")
		str := strings.Split(host, ":")
		if str[0] != "" {
			ip = str[0]
		} else {
			host = "null.com"
		}
		if len(str) > 1 {
			port, _ = strconv.Atoi(str[1])
		}

	}
	secure := false

	if r.scheme != "" {
		if r.scheme == "https" {
			secure = true
		}
	}

	if port < 1 {
		if secure {
			port = 443
		} else {
			port = 80
		}
	}

	req := &HttpStructureStandard.HttpReqData{
		Data:        message, // 请求报文
		Url:         "",
		BodyIndex:   int64(len(message) - i), // 体开始下标
		HttpVersion: http2.HTTP1,
		HttpReqService: &HttpStructureStandard.HttpReqService{
			Ip:     ip,
			Port:   int32(port),
			Secure: secure,
		},
	}

	return req
}

func BuildRequestOriginalMessage(req OriginalMessage.RequestOriginalMessage) Request {
	return Request{RequestOriginalMessage: &req}
}

// ParseRequestMessage 解析原始请求消息
func ParseRequestMessage(reqData []byte) (*Request, error) {
	message, err := OriginalMessage.ParseRequestOriginalMessage(reqData)
	if err != nil {
		return nil, err
	}
	return &Request{
		RequestOriginalMessage: &message,
	}, nil

}

func BuildRequestMessage(message OriginalMessage.RequestOriginalMessage) *Request {
	return &Request{RequestOriginalMessage: &message}
}
