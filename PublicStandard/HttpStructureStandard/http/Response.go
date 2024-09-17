package http

import (
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/OriginalMessage"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"net/http"
	"strconv"
	"strings"
)

// Response 响应
type Response struct {
	ResponseOriginalMessage *OriginalMessage.ResponseOriginalMessage // 原始请求消息
}

// GetStatusCode 获取状态码数字
func (r *Response) GetStatusCode() int {
	code := r.ResponseOriginalMessage.ResponseLine.StatusCode
	split := strings.Split(code, " ")
	statusCode, _ := strconv.Atoi(split[0])
	return statusCode
}

// GetStatusCodeString 获取状态码字符串
func (r *Response) GetStatusCodeString() string {
	return r.ResponseOriginalMessage.ResponseLine.StatusCode
}
func (r *Response) GetHttpVersion() string {
	return r.ResponseOriginalMessage.ResponseLine.HttpVersion
}

func (r *Response) GetHead() http.Header {
	return http.Header(r.ResponseOriginalMessage.Head)
}

func (r *Response) GetHeadValue(key string) []string {
	return r.ResponseOriginalMessage.Head[key]
}

func (r *Response) GetBody() []byte {
	return r.ResponseOriginalMessage.Body
}

func (r *Response) SetStatusCodeString(code string) {
	r.ResponseOriginalMessage.ResponseLine.StatusCode = code
}

func (r *Response) SetHttpVersion(version string) {
	r.ResponseOriginalMessage.ResponseLine.HttpVersion = version
}

// SetNewHead 设置新头
func (r *Response) SetNewHead(head map[string][]string) {
	r.ResponseOriginalMessage.Head = head
}

// SetHead 会覆盖原有的头键值
func (r *Response) SetHead(key string, value []string) {
	r.ResponseOriginalMessage.Head[key] = value
}

// AddHead 增加头部键值
func (r *Response) AddHead(key string, value string) {
	r.ResponseOriginalMessage.Head[key] = append(r.ResponseOriginalMessage.Head[key], value)
}

func (r *Response) SetBody(body []byte) {
	r.ResponseOriginalMessage.Body = body
}

func (r *Response) BuildMessage() []byte {
	return r.ResponseOriginalMessage.BuildMessage()
}

func (r *Response) BuildGrpcResponse() *HttpStructureStandard.HttpResData {

	message := r.BuildMessage()

	body := r.GetBody()

	res := &HttpStructureStandard.HttpResData{
		Data:        message,
		StatusCode:  int32(r.GetStatusCode()),
		BodyIndex:   int64(len(message) - len(body)),
		HttpVersion: r.GetHttpVersion(),
	}

	return res
}

func ParseResponseMessage(resData []byte) (*Response, error) {
	resMessage, err := OriginalMessage.ParseResponseOriginalMessage(resData)
	if err != nil {
		return nil, err
	}
	return &Response{ResponseOriginalMessage: &resMessage}, nil

}

func BuildResponseMessage(message OriginalMessage.ResponseOriginalMessage) *Response {
	return &Response{ResponseOriginalMessage: &message}
}
