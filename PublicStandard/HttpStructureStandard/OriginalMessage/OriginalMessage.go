package OriginalMessage

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

// http 原始报文格式

// Head 头
type Head map[string][]string

// BuildMessage 构建报文
func (h Head) BuildMessage() []byte {
	message := make([]byte, 0)
	for k, v := range h {
		for _, s := range v {
			message = append(append(message, []byte(k+": "+s)...), LineBreaks...)
		}
	}
	return message
}

// RequestOriginalMessage 请求原始报文
type RequestOriginalMessage struct {
	RequestLine RequestLine // 请求行
	Head        Head        // 头
	Body        []byte      // 请求体
}

// BuildMessage 构建报文
func (r RequestOriginalMessage) BuildMessage() []byte {
	if len(r.Body) > 0 {
		r.Head[ContentLength] = []string{strconv.Itoa(len(r.Body))} // ContentLength 只能有一个
	}

	bytes := append(r.RequestLine.BuildMessage(), r.Head.BuildMessage()...)

	bytes = append(bytes, LineBreaks...)
	bytes = append(bytes, r.Body...)
	return bytes
}

// RequestLine 请求行 由三部分组成
// 请求方法 请求路径 http版本
type RequestLine struct {
	Method      string // 请求方法
	Path        string // 请求路径包含?后的参数
	HttpVersion string // http版本 请保持使用HTTP/1.1
}

// BuildMessage 构建报文
func (l RequestLine) BuildMessage() []byte {
	return append([]byte(l.Method+" "+l.Path+" "+l.HttpVersion), LineBreaks...)
}

type Uri struct {
	Scheme     string
	Host       string
	RequestURI string
}

func (receiver Uri) BuildUrl() (*url.URL, error) {
	u := url.URL{}
	if receiver.Scheme == "" || receiver.Host == "" || receiver.RequestURI == "" {
		return nil, errors.New("scheme, host, path must not be empty")
	}
	if q := strings.Index(receiver.RequestURI, "?"); q != -1 {
		u.Path = receiver.RequestURI[:q]
		u.RawQuery = receiver.RequestURI[q+1:]
	} else {
		u.Path = receiver.RequestURI
	}
	u.Scheme = receiver.Scheme
	u.Host = receiver.Host
	return &u, nil
}

// BuildUrl 构建url
func BuildUrl(scheme, host, RequestURI string) (*url.URL, error) {
	return Uri{
		Scheme:     scheme,
		Host:       host,
		RequestURI: RequestURI,
	}.BuildUrl()
}

// ResponseOriginalMessage 响应原始报文
type ResponseOriginalMessage struct {
	ResponseLine ResponseLine // 响应行
	Head         Head         // 头
	Body         []byte       // 响应体
}

// BuildMessage 构建响应报文
func (receiver ResponseOriginalMessage) BuildMessage() []byte {
	if len(receiver.Body) > 0 {
		receiver.Head[ContentLength] = []string{strconv.Itoa(len(receiver.Body))} // ContentLength 只能有一个
	}

	bytes := append(receiver.ResponseLine.BuildMessage(), receiver.Head.BuildMessage()...)
	bytes = append(bytes, LineBreaks...)
	bytes = append(bytes, receiver.Body...)
	return bytes
}

// ResponseLine 响应行
// 由两部分组成 http版本 状态码
type ResponseLine struct {
	HttpVersion string // http版本
	StatusCode  string // 状态码 200 OK
}

func (receiver ResponseLine) BuildMessage() []byte {
	return append([]byte(receiver.HttpVersion+" "+receiver.StatusCode), LineBreaks...)
}

// ParseRequestOriginalMessage 解析请求原始报文
func ParseRequestOriginalMessage(reqData []byte) (RequestOriginalMessage, error) {

	reqMessage := RequestOriginalMessage{}
	index := 0
	for i, datum := range reqData { // 解析请求行
		if datum == Enter {
			line := ParseLine(reqData[:i])
			if len(line) != 3 {
				return RequestOriginalMessage{}, errors.New("request line format error")
			}
			reqMessage.RequestLine = RequestLine{
				Method:      line[0],
				Path:        line[1],
				HttpVersion: line[2],
			}
			index = i + 2
			break
		}
	}

	headData := reqData[index:]

	headIndex := 0

	head := make(Head)

	for i, b := range headData {
		if b == Enter { // 回车符
			if i == headIndex { // 空行
				headIndex += 2
				break
			}
			parseHead, s, err := ParseHead(headData[headIndex:i])
			if err != nil {
				return reqMessage, err
			}
			head[parseHead] = append(head[parseHead], s)
			headIndex = i + 2
		}
	}

	body := headData[headIndex:]

	reqMessage.Body = body
	reqMessage.Head = head

	return reqMessage, nil
}

// ParseResponseOriginalMessage 解析请求原始报文
func ParseResponseOriginalMessage(resData []byte) (ResponseOriginalMessage, error) {

	resMessage := ResponseOriginalMessage{}
	index := 0
	for i, datum := range resData { // 解析请求行
		if datum == Enter {
			line := ParseLine(resData[:i])
			if len(line) < 1 {
				return resMessage, errors.New("response line format error")
			}

			resMessage.ResponseLine = ResponseLine{
				HttpVersion: line[0], // http版本
				StatusCode:  strings.Join(line[1:], " "),
			}

			index = i + 2
			break
		}
	}

	headData := resData[index:]

	headIndex := 0

	head := make(Head)

	for i, b := range headData {
		if b == Enter { // 回车符
			if i == headIndex { // 空行
				headIndex += 2
				break
			}
			parseHead, s, err := ParseHead(headData[headIndex:i])
			if err != nil {
				return resMessage, err
			}
			head[parseHead] = append(head[parseHead], s)
			headIndex = i + 2
		}
	}

	body := headData[headIndex:]

	resMessage.Body = body
	resMessage.Head = head

	return resMessage, nil
}

// ParseLine 解析行 以空格分割
func ParseLine(data []byte) []string {
	return strings.Split(string(data), " ")
}

// ParseHead 解析头 仅解析一行
func ParseHead(data []byte) (string, string, error) {
	str := string(data)
	index := strings.Index(str, ": ")
	if index == -1 {
		return "", "", errors.New("head format error")
	}

	return str[:index], str[index+2:], nil
}

// NewBaseRequestOriginalMessage 创建默认请求原始报文
func NewBaseRequestOriginalMessage() RequestOriginalMessage {
	message := RequestOriginalMessage{}
	message.RequestLine = RequestLine{
		Method:      "GET",
		Path:        "/",
		HttpVersion: Http1,
	}

	head := make(Head)
	head["Host"] = []string{"mandown.xyz"}
	message.Head = head
	return message
}

// NewBaseResponseOriginalMessage 创建默认响应原始报文
func NewBaseResponseOriginalMessage() ResponseOriginalMessage {
	message := ResponseOriginalMessage{}
	message.ResponseLine = ResponseLine{
		HttpVersion: Http1,
		StatusCode:  "200 OK",
	}
	message.Head = make(Head)
	message.Head["Server"] = []string{"nginx/1.14.0"}
	message.Body = []byte("Hello World")

	return message
}
