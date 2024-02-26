package http

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type writeToData struct {
	Data []byte
}

func (w *writeToData) Write(p []byte) (n int, err error) {
	w.Data = append(w.Data, p...)
	return len(p), nil

}

//// ConvertGrpcRequest 转换为Grpc的http请求
//func ConvertGrpcRequest(request *http.Request) (*BurpMorePossibilityApi.HttpReqData, error) {
//	ww := writeToData{Data: make([]byte, 0)}
//	err := request.Write(&ww)
//	if err != nil {
//		return nil, err
//	}
//
//	secure := false
//	port := 0
//	if request.URL.Scheme == "StandardHttp" {
//		secure = false
//		port = 80
//	}
//	if request.URL.Scheme == "https" {
//		secure = true
//		port = 443
//	}
//	if request.URL.Port() != "" {
//		port, err = strconv.Atoi(request.URL.Port())
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	reqTest := &BurpMorePossibilityApi.HttpReqData{
//		Data:        ww.Data,
//		Url:         request.URL.String(),
//		BodyIndex:   int64(len(ww.Data)) - request.ContentLength,
//		HttpVersion: request.Proto,
//		HttpReqService: &BurpMorePossibilityApi.HttpReqService{
//			Ip:     request.Host,
//			Port:   int32(port),
//			Secure: secure,
//		},
//	}
//	return reqTest, nil
//}
//
//// ConvertGrpcResponse 转换为Grpc的http响应
//func ConvertGrpcResponse(res *http.Response) (*BurpMorePossibilityApi.HttpResData, error) {
//	head := writeToData{Data: make([]byte, 0)}
//
//	body, err := io.ReadAll(res.Body)
//	if err != nil {
//		return nil, err
//	}
//	res.Header.Set("Content-Length", strconv.Itoa(len(body)))
//
//	err = res.Header.Write(&head)
//	if err != nil {
//		return nil, err
//	}
//	resDate := []byte(res.Proto + " " + res.Status + "\r\n")
//	resDate = append(resDate, head.Data...)
//	resDate = append(resDate, []byte("\r\n")...)
//	resDate = append(resDate, body...)
//
//	return &BurpMorePossibilityApi.HttpResData{
//		Data:        resDate,
//		StatusCode:  int32(res.StatusCode),
//		BodyIndex:   int64(len(resDate) - len(body)),
//		HttpVersion: res.Proto,
//	}, nil
//}

//// testModifier 测试修饰器
//func (h HttpToolProcess) testModifier() {
//	//h.Modifier
//}

//// ConvertReqStandardHttp 请求转换为标准http组
//func ConvertReqStandardHttp(req *BurpMorePossibilityApi.HttpReqData) (StandardHttp, error) {
//	s := StandardHttp{}
//	c := s.ConvertHttpReqOwn(req)
//	return s, c
//}
//
//// ConvertResStandardHttp 响应转换为标准http组
//func ConvertResStandardHttp(res *BurpMorePossibilityApi.HttpResData) (StandardHttp, error) {
//	s := StandardHttp{}
//	c := s.ConvertHttpResOwn(res)
//	return s, c
//}
//
//// ConvertStandardHttp 转换为标准http组
//func ConvertStandardHttp(req *BurpMorePossibilityApi.HttpReqData, res *BurpMorePossibilityApi.HttpResData) (StandardHttp, error) {
//	s := StandardHttp{}
//	c := s.ConvertHttpReqOwn(req)
//	if c != nil {
//		return s, c
//	}
//	c = s.ConvertHttpResOwn(res)
//	return s, c
//}

//// StandardHttp 标准http 基于标准库中的http封装
//type StandardHttp struct {
//	Req *StandardHttpReq // 请求
//	Res *StandardHttpRes // 响应
//}

//// ConvertHttpReqDate 转为burpAPI中的请求类型
//func (s *StandardHttp) ConvertHttpReqDate() (*BurpMorePossibilityApi.HttpReqData, error) {
//	req := s.Req
//	if req == nil {
//		return nil, fmt.Errorf("StandardHttp request is empty and cannot be converted")
//	}
//	ww := writeToData{Data: make([]byte, 0)}
//
//	head := req.ReqDate[:req.BodyIndex]
//
//	index := strings.Index(string(head), "\r\n")
//	up := []byte(string(head)[:index+2])
//
//	//up := []byte(req.Method + " " + req.Url.RawPath + "HTTP/1.1")
//	req.ReqHead.Set("Content-Length", strconv.Itoa(len(req.ReqBody)))
//	err := req.ReqHead.Write(&ww)
//	if err != nil {
//		return nil, err
//	}
//	heads := ww.Data
//	up = append(up, heads...)
//	up = append(up, []byte("\r\n")...)
//	up = append(up, req.ReqBody...)
//	return &BurpMorePossibilityApi.HttpReqData{
//		Data:           up,
//		Url:            req.Url.String(),
//		BodyIndex:      int64(len(up) - len(req.ReqBody)),
//		HttpVersion:    req.HttpVersion,
//		HttpReqService: req.TarGetInfo,
//	}, nil
//}
//
//// ConvertHttpReqOwn 将grpc的http请求转为标准请求
//func (s *StandardHttp) ConvertHttpReqOwn(data *BurpMorePossibilityApi.HttpReqData) error {
//	d := data.GetData()
//	request, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(CompelReqHttp1(data))))
//	if err != nil {
//		return err
//	}
//	parse, err := url.Parse(data.GetUrl())
//	if err != nil {
//		return err
//	}
//
//	body := d[data.GetBodyIndex():]
//
//	s.Req = &StandardHttpReq{ // 构建标准请求
//		Url:         parse,
//		TarGetInfo:  data.GetHttpReqService(),
//		ReqHead:     request.Header,
//		ReqBody:     body,
//		Method:      request.Method,
//		HttpVersion: data.HttpVersion,
//		ReqDate:     d,
//		BodyIndex:   data.GetBodyIndex(),
//	}
//
//	return nil
//}
//
//func (s *StandardHttp) ConvertHttpResDate() (*BurpMorePossibilityApi.HttpResData, error) {
//	if s.Res == nil {
//		return nil, fmt.Errorf("StandardHttp response is empty and cannot be converted")
//	}
//	res := *s.Res
//	return &BurpMorePossibilityApi.HttpResData{
//		Data:        res.ResDate,
//		StatusCode:  res.StatusCode,
//		BodyIndex:   res.BodyIndex,
//		HttpVersion: res.HttpVersion,
//	}, nil
//}
//
//// ConvertHttpResOwn 将grpc的http响应转回标准http组
//func (s *StandardHttp) ConvertHttpResOwn(data *BurpMorePossibilityApi.HttpResData) error {
//	resData := data.GetData()
//	response, err := http.ReadResponse(bufio.NewReader(bytes.NewReader(CompelResHttp1(data))), nil)
//	if err != nil {
//		return err
//	}
//	s.Res = &StandardHttpRes{
//		StatusCode:  int32(response.StatusCode),
//		ResHead:     response.Header,
//		ResBody:     resData[data.GetBodyIndex():],
//		HttpVersion: data.HttpVersion,
//		ResDate:     resData,
//		BodyIndex:   data.GetBodyIndex(),
//	}
//	return nil
//}
//
//// BuildRequest 构建http.Request
//func (s *StandardHttp) BuildRequest() (*http.Request, error) {
//	request, err := http.NewRequest(s.Req.Method, s.Req.Url.String(), bytes.NewReader(s.Req.ReqBody))
//	if err != nil {
//		return nil, err
//	}
//	request.Header = s.Req.ReqHead
//	return request, nil
//}

//// StandardHttpReq 标准请求
//type StandardHttpReq struct {
//	Url         *url.URL
//	TarGetInfo  *BurpMorePossibilityApi.HttpReqService // 目标信息
//	ReqHead     http.Header
//	ReqBody     []byte // 请求体
//	Method      string // 请求方法
//	HttpVersion string // http版本
//	ReqDate     []byte // 完整请求报文 指针禁止修改
//	BodyIndex   int64  // 体开始下标 用于后续截取请求体并计算长度
//}
//
//// StandardHttpRes 标准响应
//type StandardHttpRes struct {
//	StatusCode  int32       // 响应码
//	ResHead     http.Header // 响应头
//	ResBody     []byte      // 响应体
//	HttpVersion string      // http版本
//	ResDate     []byte      // 完整响应报文 指针禁止修改
//	BodyIndex   int64       // 体开始下标 用于后续截取请求体并计算长度
//}

// HttpStandardGroup 标准http组
type HttpStandardGroup struct {
	Req *RefactorStandardHttpReq // 请求
	Res *RefactorStandardHttpRes // 响应
}

// CompelReqHttp1 强制http1避免构造出错,客户端会自动处理通信时http2的处理
func CompelReqHttp1(r BurpHttpData) []byte {
	httpData := r.GetData()                     // 请求报文
	head := string(httpData[:r.GetBodyIndex()]) // 头
	body := httpData[r.GetBodyIndex():]         // 体
	index := strings.Index(head, "\r\n")        // 检索第一行
	if index == -1 {
		return nil
	}
	one := head[:index]                        // 第一行
	httpVersion := strings.LastIndex(one, " ") // 第一行的倒数第一个空格之后就是http的版本
	o := one[:httpVersion]                     //
	o += " HTTP/1.1"
	o += head[index:]
	by := []byte(o)
	return append(by, body...)
}

// CompelResHttp1 强制http1避免构造出错,客户端会自动处理通信时http2的处理
func CompelResHttp1(r BurpHttpData) []byte {
	head := string(r.GetData()[:r.GetBodyIndex()]) // 头
	index := strings.Index(head, "\r\n")           // 检索第一行
	if index == -1 {
		return nil
	}
	up := head[:index+2] // 第一行
	i := strings.Index(up, " ")
	newUp := HTTP1 + up[i:]
	newUp += head[index+2:]
	return append([]byte(newUp), r.GetData()[r.GetBodyIndex():]...)

}

// RefactorStandardHttpRes 重构标准响应
type RefactorStandardHttpRes struct {
	statusCode   int32       // 响应码
	codeString   string      // 响应码字符串 200 OK 取后面的OK
	resHead      http.Header // 响应头
	resBody      []byte      // 响应体
	httpVersion  string      // http版本
	rawResDate   []byte      // 原始响应报文
	rawBodyIndex int64       // 体开始下标 用于后续截取请求体并计算长度

	standardHttpReq *RefactorStandardHttpReq // 请求

	//upRow []byte // 头行

}

// GetReq 获取请求
func (r *RefactorStandardHttpRes) GetReq() *RefactorStandardHttpReq {
	return r.standardHttpReq
}

// GetStatusCode 获取响应码
func (r *RefactorStandardHttpRes) GetStatusCode() int32 {
	return r.statusCode
}

// GetStatusCodeString 获取响应码字符串 200 OK 取后面的OK
func (r *RefactorStandardHttpRes) GetStatusCodeString() string {
	return r.codeString
}

// GetResHead 获取响应头 为引用类型修改后此实例头也会跟着改动
func (r *RefactorStandardHttpRes) GetResHead() http.Header {
	return r.resHead
}

// GetResBody 获取响应体
func (r *RefactorStandardHttpRes) GetResBody() []byte {
	return r.resBody
}

// GetHttpVersion 获取http版本
func (r *RefactorStandardHttpRes) GetHttpVersion() string {
	return r.httpVersion
}

// GetRawResData 获取原始响应报文
func (r *RefactorStandardHttpRes) GetRawResData() []byte {
	return r.rawResDate
}

// GetRawBodyIndex 获取原始响应体开始下标
func (r *RefactorStandardHttpRes) GetRawBodyIndex() int64 {
	return r.rawBodyIndex
}

// SetStatusCode 设置响应码
func (r *RefactorStandardHttpRes) SetStatusCode(code int32) {
	r.statusCode = code
}

// SetStatusCodeString 设置响应码字符串 200 OK 包含状态码和状态码字符串
func (r *RefactorStandardHttpRes) SetStatusCodeString(code string) error {
	index := strings.Index(code, " ")
	if index == -1 {
		return fmt.Errorf("status code string is empty")
	}
	r.codeString = code[index+1:]
	statusCode, err := strconv.Atoi(code[:index])
	if err != nil {
		return err
	}
	r.statusCode = int32(statusCode)
	return nil
}

// SetNewHeader 设置新头
func (r *RefactorStandardHttpRes) SetNewHeader(head http.Header) {
	r.resHead = head
}

// SetHttpVersion 设置http版本 警告: 请勿随意修改 需要和请求的http版本一致
// 除非你知道你在做什么
func (r *RefactorStandardHttpRes) SetHttpVersion(httpVersion string) {
	r.httpVersion = httpVersion
}

// SetResBody 设置响应体
func (r *RefactorStandardHttpRes) SetResBody(body []byte) {
	//h := r.rawResDate[:r.rawBodyIndex]         // 头
	//h = append(h, body...)                  // 合并响应报文
	//r.rawBodyIndex = int64(len(h) - len(body)) // 体开始下标
	r.resBody = body
	//r.rawResDate = h
	//r.resHead.Set(standard.ContentLength, strconv.Itoa(len(body))) // 设置ContentLength 长度
}

// BuildGrpcRes 构建Grpc的响应类型
func (r *RefactorStandardHttpRes) BuildGrpcRes() (*HttpStructureStandard.HttpResData, error) {
	r.resHead.Set(ContentLength, strconv.Itoa(len(r.resBody)))                                        // 设置ContentLength 长度
	up := []byte(r.httpVersion + " " + strconv.Itoa(int(r.statusCode)) + " " + r.codeString + "\r\n") // 头行
	ww := writeToData{Data: make([]byte, 0)}
	err := r.resHead.Write(&ww) // 读响应头
	if err != nil {
		return nil, err
	}
	// 组装完整响应报文
	up = append(up, ww.Data...)
	up = append(up, []byte("\r\n")...)
	up = append(up, r.resBody...)
	return &HttpStructureStandard.HttpResData{
		Data:        up,
		StatusCode:  r.statusCode,
		BodyIndex:   int64(len(up) - len(r.resBody)),
		HttpVersion: r.httpVersion,
	}, nil
}

// BuildResponse 构建响应
func (r *RefactorStandardHttpRes) BuildResponse() (*http.Response, error) {
	res, err := r.BuildGrpcRes() // 先构建grpc的响应
	if err != nil {
		return nil, err
	}
	return http.ReadResponse(bufio.NewReader(bytes.NewReader(CompelResHttp1(res))), nil)
}

// BuildHttpReqAndRes 构建http请求和响应
func (r *RefactorStandardHttpRes) BuildHttpReqAndRes() (*HttpStructureStandard.HttpReqAndRes, error) {
	req, err := r.standardHttpReq.BuildReqData()
	res, err := r.BuildGrpcRes()
	if err != nil {
		return nil, err
	}
	return &HttpStructureStandard.HttpReqAndRes{
		Req: req,
		Res: res,
	}, nil
}

// ConvertHttpResOwn 转换为自身响应
func (r *RefactorStandardHttpRes) ConvertHttpResOwn(res *HttpStructureStandard.HttpResData, req *RefactorStandardHttpReq) error {

	response, err := BuildResponse(res)
	if err != nil {
		return err
	}
	r.standardHttpReq = req
	//http1Res := CompelResHttp1(res)
	r.statusCode = res.GetStatusCode()

	index := strings.Index(response.Status, " ")
	r.codeString = response.Status[index+1:] // 响应码字符串 200 OK 取后面的OK

	//r.upRow = []byte(response.Proto + " " + response.Status + "\r\n") // 头行

	r.resHead = response.Header
	r.resBody = res.GetData()[res.GetBodyIndex():]
	r.httpVersion = res.GetHttpVersion()
	r.rawResDate = res.GetData()
	r.rawBodyIndex = res.GetBodyIndex()
	return nil
}

// ConvertHttpResponseOwn 转换为自身响应
func (r *RefactorStandardHttpRes) ConvertHttpResponseOwn(response *http.Response, req *RefactorStandardHttpReq) error {
	r.standardHttpReq = req
	r.statusCode = int32(response.StatusCode)
	index := strings.Index(response.Status, " ")
	r.codeString = response.Status[index+1:] // 响应码字符串 200 OK 取后面的OK
	r.resHead = response.Header
	body, err := io.ReadAll(response.Body)

	if err != nil {
		if body == nil {
			return err
		}
	}
	r.resBody = body
	r.httpVersion = response.Proto
	return nil
}

//func (r *RefactorStandardHttpRes) TestRes() {
//
//	fmt.Println(r.resHead)
//	fmt.Println(r.statusCode)
//	fmt.Println(r.httpVersion)
//
//}

// RefactorStandardHttpReq 重构标准请求
type RefactorStandardHttpReq struct {
	httpReqService *HttpStructureStandard.HttpReqService // 目标信息
	httpVersion    string                                // http版本
	request        *http.Request                         // 请求
	body           []byte                                // 体
	client         HttpClient                            // 客户端
}

// SetClient 设置客户端
func (r *RefactorStandardHttpReq) SetClient(client HttpClient) {
	r.client = client
}

// GetClient 获取客户端
func (r *RefactorStandardHttpReq) GetClient() HttpClient {
	return r.client
}

// GetUrl 获取完整URL 包含协议头和地址
func (r *RefactorStandardHttpReq) GetUrl() string {
	return r.request.URL.String()
}

// GetRawPath 获取请求路径以及参数 不包含协议头和地址,
// 参数解析可以用 url.ParseQuery
func (r *RefactorStandardHttpReq) GetRawPath() string {
	return UrlToRawPath(r.request.URL)
}

// GetTarGetPath 获取目标地址
func (r *RefactorStandardHttpReq) GetTarGetPath() *HttpStructureStandard.HttpReqService {
	return r.httpReqService
}

// GetPath 获取请求路径 不包含参数
func (r *RefactorStandardHttpReq) GetPath() string {
	return r.request.URL.Path
}

// GetBody 获取请求体
func (r *RefactorStandardHttpReq) GetBody() []byte {
	return r.body
}

// GetHead 获取头 为引用类型修改后此实例头也会跟着改动
func (r *RefactorStandardHttpReq) GetHead() http.Header {
	return r.request.Header
}

// GetMethod 获取请求方法
func (r *RefactorStandardHttpReq) GetMethod() string {
	return r.request.Method
}

// SetTarGetPath 设置目标地址
func (r *RefactorStandardHttpReq) SetTarGetPath(tarGet *HttpStructureStandard.HttpReqService) {
	//r.request.URL.
	r.request.URL.Host = fmt.Sprintf("%s:%d", tarGet.GetIp(), tarGet.GetPort())
	if tarGet.GetSecure() {
		r.request.URL.Scheme = "https"
	} else {
		r.request.URL.Scheme = "http"
	}
	r.httpReqService = tarGet
}

// SetHead 设置新头
func (r *RefactorStandardHttpReq) SetHead(head http.Header) {
	r.request.Header = head
}

// SetHostHerder 设置Host头
func (r *RefactorStandardHttpReq) SetHostHerder(host string) {
	r.request.Host = host
	r.request.Header.Set("Host", host)
}

// SetRawPath 设置路径 包含参数 设置完整的Path 会覆盖原有的参数
func (r *RefactorStandardHttpReq) SetRawPath(rawPath string) error {
	parse, err := url.Parse(rawPath)
	if err != nil {
		return err
	}
	r.request.URL.Path = parse.Path
	r.request.URL.RawQuery = parse.RawQuery
	return nil
	//r.request.URL.
}

// SetPath 设置请求路径 参数不变
func (r *RefactorStandardHttpReq) SetPath(path string) {
	r.request.URL.Path = path
}

// SetQuery 设置请求参数
func (r *RefactorStandardHttpReq) SetQuery(query string) {
	r.request.URL.RawQuery = query
}

// SetUrl 设置url
func (r *RefactorStandardHttpReq) SetUrl(uri string) error {
	parse, err := url.Parse(uri)
	if err != nil {
		return err
	}
	r.request.URL = parse
	return nil
}

// SetMethod 设置请求方法
func (r *RefactorStandardHttpReq) SetMethod(method string) {
	r.request.Method = method
}

// SetBody 设置请求体
func (r *RefactorStandardHttpReq) SetBody(body []byte) {
	r.body = body
}

// ConvertHttpReqOwn 转换为自身请求
func (r *RefactorStandardHttpReq) ConvertHttpReqOwn(req GrpcHttpReq) error {
	//http1 := CompelReqHttp1(req) // 强转为http1

	request, err := BuildRequest(req)
	if err != nil {
		return err
	}
	r.request = request
	uri, err := url.Parse(req.GetUrl())
	if err != nil {
		return err
	}
	request.URL = uri
	r.SetHostHerder(uri.Host)
	//r.url = uri
	r.body = req.GetData()[req.GetBodyIndex():]
	r.httpVersion = req.GetHttpVersion()
	r.httpReqService = req.GetHttpReqService()
	return nil
}

// ConvertHttpRequestOwn 转换为自身请求
func (r *RefactorStandardHttpReq) ConvertHttpRequestOwn(req *http.Request) error {

	r.request = req
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	r.body = body
	r.httpVersion = req.Proto

	return nil
}

// BuildReqData 构建Grpc的请求类型
func (r *RefactorStandardHttpReq) BuildReqData() (*HttpStructureStandard.HttpReqData, error) {
	r.request.Header.Set(ContentLength, strconv.Itoa(len(r.body))) // 设置ContentLength 长度
	r.request.ContentLength = int64(len(r.body))
	up := []byte(r.request.Method + " " + UrlToRawPath(r.request.URL) + " " + HTTP1 + "\r\n")
	ww := writeToData{Data: make([]byte, 0)}
	err := r.request.Header.Write(&ww)
	if err != nil {
		return nil, err
	}
	// 组装完整请求报文
	up = append(up, ww.Data...)
	up = append(up, []byte("\r\n")...)
	up = append(up, r.body...)
	return &HttpStructureStandard.HttpReqData{
		Data:           up,
		Url:            r.request.URL.String(),
		BodyIndex:      int64(len(up) - len(r.body)), // 体开始下标
		HttpVersion:    r.httpVersion,
		HttpReqService: r.httpReqService,
	}, nil
}

// BuildRequest 构建请求
func (r *RefactorStandardHttpReq) BuildRequest() *http.Request {
	r.request.Header.Set(ContentLength, strconv.Itoa(len(r.body))) // 设置ContentLength 长度
	r.request.ContentLength = int64(len(r.body))
	r.request.Body = io.NopCloser(bytes.NewReader(r.body))
	r.request.RequestURI = ""
	return r.request
}

// Send 发送请求
func (r *RefactorStandardHttpReq) Send() (standardHttpRes RefactorStandardHttpRes, err error) {
	res, err := r.client.Do(r.BuildRequest())
	if err != nil {
		return RefactorStandardHttpRes{}, err
	}
	return BuildRefactorStandardHttpResponse(res, r)
}

// BuildRefactorStandardHttpRes	构建重构标准响应 用BurpApi.HttpResData 构建
func BuildRefactorStandardHttpRes(res *HttpStructureStandard.HttpResData, req *RefactorStandardHttpReq) (RefactorStandardHttpRes, error) {
	httpRes := RefactorStandardHttpRes{}
	err := httpRes.ConvertHttpResOwn(res, req)
	if err != nil {
		return RefactorStandardHttpRes{}, err
	}
	return httpRes, nil
}

// BuildRefactorStandardHttpResponse 构建重构标准响应 用http.Response 构建
func BuildRefactorStandardHttpResponse(res *http.Response, req *RefactorStandardHttpReq) (RefactorStandardHttpRes, error) {
	httpRes := RefactorStandardHttpRes{}
	err := httpRes.ConvertHttpResponseOwn(res, req)
	if err != nil {
		return RefactorStandardHttpRes{}, err
	}
	return httpRes, nil
}

// BuildRefactorStandardHttpReq 构建转换为自身请求
func BuildRefactorStandardHttpReq(req *HttpStructureStandard.HttpReqData, client HttpClient) (RefactorStandardHttpReq, error) {
	httpReq := RefactorStandardHttpReq{}
	httpReq.SetClient(client)
	return httpReq, httpReq.ConvertHttpReqOwn(req)
}

// BuildRequest 构建请求 将grpc的请求转换为http请求
func BuildRequest(req GrpcHttpReq) (*http.Request, error) {
	return http.ReadRequest(bufio.NewReader(bytes.NewReader(CompelReqHttp1(req))))
}

// BuildResponse 构建响应 将grpc的响应转换为http响应
func BuildResponse(res GrpcHttpRes) (*http.Response, error) {
	return http.ReadResponse(bufio.NewReader(bytes.NewReader(CompelResHttp1(res))), nil)
}

// BuildGrpcRequest 构建Grpc请求实例  <===== 等待测试 =====>
func BuildGrpcRequest(req *http.Request) (*HttpStructureStandard.HttpReqData, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		body = []byte{}
	}
	req.Header.Set(ContentLength, strconv.Itoa(len(body))) // 设置ContentLength 长度
	req.ContentLength = int64(len(body))
	up := []byte(req.Method + " " + UrlToRawPath(req.URL) + " " + HTTP1 + "\r\n")
	ww := writeToData{Data: make([]byte, 0)}
	err = req.Header.Write(&ww)
	if err != nil {
		return nil, err
	}
	// 组装完整请求报文
	up = append(up, ww.Data...)
	up = append(up, []byte("\r\n")...)
	up = append(up, body...)

	atoi, err := strconv.Atoi(req.URL.Port())
	if err != nil {
		return nil, err
	}

	Secure := false
	if strings.Contains(req.URL.Scheme, "s") {
		Secure = true
	}
	return &HttpStructureStandard.HttpReqData{
		Data:        up,
		Url:         req.URL.String(),
		BodyIndex:   int64(len(up) - len(body)), // 体开始下标
		HttpVersion: req.Proto,
		HttpReqService: &HttpStructureStandard.HttpReqService{
			Ip:     req.URL.Hostname(),
			Port:   int32(atoi),
			Secure: Secure,
		},
	}, nil
}

// BuildRefactorStandardHttpRequest 构建转换为自身请求
func BuildRefactorStandardHttpRequest(req *http.Request, client HttpClient) (RefactorStandardHttpReq, error) {
	httpReq := RefactorStandardHttpReq{}
	httpReq.SetClient(client)
	return httpReq, httpReq.ConvertHttpRequestOwn(req)
}

// BuildStandardHttpGroup 构建标准http组
func BuildStandardHttpGroup(httpGroup *HttpStructureStandard.HttpReqAndRes, client HttpClient) (HttpStandardGroup, error) {
	group := HttpStandardGroup{}
	req := httpGroup.GetReq()
	own, err := BuildRefactorStandardHttpReq(req, client)
	if err != nil {
		return group, err
	}

	group.Req = &own
	res, err := BuildRefactorStandardHttpRes(httpGroup.GetRes(), &own)
	if err != nil {
		return group, err
	}
	group.Res = &res

	return group, nil
}

// UrlToRawPath 拿到请求路劲的部分包含参数
func UrlToRawPath(r *url.URL) string {
	q := ""
	if r.RawQuery != "" {
		q = "?" + r.RawQuery
	}
	if r.RawPath != "" {
		return r.RawPath + q
	}
	if r.Path != "" {
		return r.Path + q
	}
	return "/"
}

// NewRefactorStandardHttpReq 新建标准请求
func NewRefactorStandardHttpReq(url string, method string, body []byte, client HttpClient) (RefactorStandardHttpReq, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return RefactorStandardHttpReq{}, err
	}
	req.Header = http.Header{
		"User-Agent": []string{DefaultUA},
	}
	return BuildRefactorStandardHttpRequest(req, client)
}
