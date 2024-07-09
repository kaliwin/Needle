package OriginalMessage

var LineBreaks = []byte{Enter, Line} // 回车换行 \r\n

const Enter = 13
const Line = 10

const (
	Http1 = "HTTP/1.1"
	Http2 = "HTTP/2"
)

// 标准下该有的头部

const (
	Host          = "Host"           // 请求主机
	UA            = "User-Agent"     // UA
	ContentLength = "Content-Length" // 体长度
	Accept        = "Accept"         // 接受类型   */* 表示任意类型
)
