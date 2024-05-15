package sign

type SignType string // 签名类型

const (
	UrlSignType       SignType = "UrlSignType"       // url 签名 正常情况下不需要计算请求体
	BodySignType      SignType = "BodySignType"      // 体签名 字节流的签名
	UuidSignType      SignType = "UuidSignType"      // uuid 唯一标识符
	HttpGroupSignType SignType = "HttpGroupSignType" // http组签名
)
