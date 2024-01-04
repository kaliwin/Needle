package HttpStructureStandard

// 标准http结构  ,所有程序都应该使用这个结构
// 不同http库的结构都应该可以转换为这个结构 , 并且可以从这个结构转换为其他http库的结构

// es、bleve 等搜索引擎也要使用这个结构

// HttpRequest http请求结构
type HttpRequest struct {
	// 请求路径包含参数
	RawPath string `json:"rawPath"`
	// 请求方法
	Method string `json:"method"`
	// 请求头
	Header map[string][]string `json:"header"`
	// 请求体
	Body []byte `json:"body"`
	// 目标地址以及通信协议
	TarGetAddress TarGetAddress `json:"tarGetAddress"`
}

// TarGetAddress 目标地址以及协议 请求报文中表示
type TarGetAddress struct {
	Host   string `json:"host"`   // 主机 ip、域名
	Port   int    `json:"port"`   // 端口
	Scheme string `json:"scheme"` // 协议
}

// HttpResponse http响应结构
type HttpResponse struct {
	HttpVersion string              `json:"httpVersion"` // http协议版本
	StatusCode  string              `json:"statusCode"`  // 状态码
	Header      map[string][]string `json:"header"`      // 响应头
	Body        []byte              `json:"body"`        // 响应体
}
