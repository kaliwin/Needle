package http

import (
	"bufio"
	"bytes"
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
	uri                    *Uri
}

// GetScheme 获取协议
func (r *Request) GetScheme() string {
	return r.uri.GetScheme()
}

// SetScheme 设置协议
func (r *Request) SetScheme(Scheme string) {
	r.uri.SetScheme(Scheme)
}

// GetUrl 获取url
func (r *Request) GetUrl() *url.URL {
	return r.uri.BuildUrl()
}

// SetUrl 设置url
func (r *Request) SetUrl(url *url.URL) {
	r.uri.uri = url
}

// GetMethod 获取请求方法
func (r *Request) GetMethod() string {
	return r.RequestOriginalMessage.RequestLine.Method
}

// SetMethod 设置请求方法
func (r *Request) SetMethod(method string) {
	r.RequestOriginalMessage.RequestLine.Method = method
}

// GetPath 获取请求路径不包含?后的参数
func (r *Request) GetPath() string {
	return r.uri.GetPath()
}

// SetPath 设置请求路径 不包含参数
func (r *Request) SetPath(path string) {
	r.uri.SetPath(path)
}

func (r *Request) AddPath(path string) {
	r.uri.IntelligentAddPath(path)
}

func (r *Request) DeleteQuery(key string) {

	r.uri.DeleteQuery(key)

}

// GetQuery 获取参数
func (r *Request) GetQuery() url.Values {
	return r.uri.query
}

// SetQuery 设置参数
func (r *Request) SetQuery(q url.Values) {
	r.uri.query = q
}

// GetQueryValue 获取参数值
func (r *Request) GetQueryValue(key string) []string {
	return r.uri.GetQueryValue(key)
}

// SetQueryValue 设置参数 会覆盖原先的
func (r *Request) SetQueryValue(key, value string) {
	r.uri.SetQueryValue(key, value)
}

// AddQueryValue 添加参数
func (r *Request) AddQueryValue(key, value string) {
	r.uri.AddQueryValue(key, value)
}

func (r *Request) DeleteHead(key string) {
	delete(r.RequestOriginalMessage.Head, key)
}

// GetHead 获取头部实例
func (r *Request) GetHead() http.Header {
	return http.Header(r.RequestOriginalMessage.Head)
}

// GetHeadValue 获取头部值
func (r *Request) GetHeadValue(key string) []string {
	return r.RequestOriginalMessage.Head[key]
}

// SetNewHead 设置新的头部
func (r *Request) SetNewHead(head map[string][]string) {
	r.RequestOriginalMessage.Head = head
}

// SetHead 会覆盖原有的头键值
func (r *Request) SetHead(key string, value []string) {
	r.RequestOriginalMessage.Head[key] = value
}

func (r *Request) SetHeadValue(key, value string) {
	r.RequestOriginalMessage.Head[key] = []string{value}
}

// AddHead 增加头部键值
func (r *Request) AddHead(key string, value string) {
	r.RequestOriginalMessage.Head[key] = append(r.RequestOriginalMessage.Head[key], value)
}

// SetBody 设置请求体
func (r *Request) SetBody(body []byte) {
	r.RequestOriginalMessage.Body = body
}

// GetBody 获取请求体
func (r *Request) GetBody() []byte {
	return r.RequestOriginalMessage.Body
}

// BuildMessage 构建请求报文
func (r *Request) BuildMessage() []byte {
	r.RequestOriginalMessage.RequestLine.Path = r.uri.BuildPathLine()
	return r.RequestOriginalMessage.BuildMessage()
}

// BuildRawHttpRequest 构建原始请求 没有处理URL
// 强制使用HTTP/1.1
func (r *Request) BuildRawHttpRequest() (*http.Request, error) {
	r.RequestOriginalMessage.RequestLine.HttpVersion = http2.HTTP1

	request, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(r.BuildMessage())))

	if host := r.GetHeadValue("Host"); len(host) != 0 {
		request.Host = host[0]
	}

	if err != nil {
		return nil, err
	}
	return request, err
}

//// BuildHttpRequest 构建可发送的请求 没有设置url 会从原始请求中获取到Host确定目标 但是需要指定协议
//// url 不为空 会直接使用url
//func (r *Request) BuildHttpRequest() (*http.Request, error) {
//	request, err := r.BuildRawHttpRequest()
//	if err != nil {
//		return request, err
//	}
//	request.URL = r.uri.BuildUrl()
//	request.RequestURI = ""
//	return request, nil
//
//}

// BuildGrpcRequest 构建grpc请求实例
func (r *Request) BuildGrpcRequest() *HttpStructureStandard.HttpReqData {

	message := r.BuildMessage()
	i := len(r.GetBody())

	ip := ""
	port := 0

	if r.uri.GetHost() != "" {
		str := strings.Split(r.uri.GetHost(), ":")
		if str[0] != "" {
			ip = str[0]
		}
		if len(str) > 1 {
			port, _ = strconv.Atoi(str[1])
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

	if r.GetScheme() != "" {
		if r.GetScheme() == "https" {
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
		Url:         r.uri.BuildUrl().String(),
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

	host := ""

	if len(message.Head["Host"]) != 0 {
		host = message.Head["Host"][0]
	}

	parse, err := url.Parse(message.RequestLine.Path)
	if err != nil {
		return nil, err
	}

	parse.Host = host

	return &Request{
		RequestOriginalMessage: &message,
		uri: &Uri{
			uri:   parse,
			query: parse.Query(),
		},
	}, nil

}

func BuildRequestMessage(message OriginalMessage.RequestOriginalMessage) *Request {
	return &Request{RequestOriginalMessage: &message}
}

func ParseGrpcRequest(req *HttpStructureStandard.HttpReqData) (*Request, error) {
	message, err := ParseRequestMessage(req.GetData())
	if err != nil {
		return message, err
	}
	parse, err := url.Parse(req.GetUrl())
	if err != nil {
		return message, err
	}
	uri := NewUri(parse)
	message.uri = uri
	return message, nil
}

func ParseHttpRequest(reqData *http.Request) (*Request, error) {
	all := ReadAll{}
	err := reqData.Write(&all)
	if err != nil {
		return nil, err
	}
	message, err := ParseRequestMessage(all.Data)
	if err != nil {
		return nil, err
	}
	if reqData.URL != nil {
		message.SetUrl(reqData.URL)
	}

	return message, nil
}

type ReadAll struct {
	Data []byte
}

func (r *ReadAll) Write(p []byte) (n int, err error) {
	r.Data = append(r.Data, p...)
	return len(p), nil
}

func NewUri(uri *url.URL) *Uri {
	return &Uri{
		uri:   uri,
		query: uri.Query(),
	}
}

type Uri struct {
	uri   *url.URL
	query url.Values
}

func (u *Uri) GetScheme() string {
	return u.uri.Scheme
}

func (u *Uri) SetScheme(scheme string) {
	u.uri.Scheme = scheme
}

// GetQuery 获取参数
func (u *Uri) GetQuery() url.Values {
	return u.query
}

// SetQuery 设置参数
func (u *Uri) SetQuery(q url.Values) {
	u.query = q
}
func (u *Uri) DeleteQuery(key string) {
	u.query.Del(key)
}

// GetQueryValue 获取参数值
func (u *Uri) GetQueryValue(key string) []string {
	return u.query[key]
}

// SetQueryValue 设置参数 会覆盖原先的
func (u *Uri) SetQueryValue(key, value string) {
	u.query.Set(key, value)
}

// AddQueryValue 添加参数
func (u *Uri) AddQueryValue(key, value string) {
	u.query.Add(key, value)
}

// GetPath 获取路径
func (u *Uri) GetPath() string {
	return u.uri.Path
}

// SetPath 设置路径覆盖原先的
func (u *Uri) SetPath(path string) {
	u.uri.Path = path
}

// AddPath 添加路径 会拼接上原先的
func (u *Uri) AddPath(path string) {
	u.uri.Path += path
}

// IntelligentAddPath 智能添加路径 确保不能出现两个连着的//
func (u *Uri) IntelligentAddPath(path string) {
	p := u.uri.Path

	if p != "" {
		if strings.LastIndex(p, "/")+1 == len(p) {
			p = p[:len(p)-1]
		}
	}

	if strings.Index(path, "/") == 0 {
		path = path[1:]
	}
	u.uri.Path = p + "/" + path
}

func (u *Uri) GetHost() string {
	return u.uri.Host
}

// SetHost 设置目标地址需要带上端口 www.baidu.com:443
func (u *Uri) SetHost(host string) {
	u.uri.Host = host
}

func (u *Uri) BuildPathLine() string {
	buildUrl := u.BuildUrl()
	path := buildUrl.Path
	if path == "" {
		path = "/"
	}
	if buildUrl.RawQuery != "" {
		path += "?" + buildUrl.RawQuery
	}
	return path
}

func (u *Uri) BuildUrl() *url.URL {
	u.uri.RawQuery = u.query.Encode()
	if u.uri.Path == "" {
		u.uri.Path = "/"
	}
	return u.uri
}
