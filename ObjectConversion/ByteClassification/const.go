package ByteClassification

type ByteType string // 字节类型

const (
	StringByteType  ByteType = "String"  // 字符串类型 UTF-8
	UnknownByteType ByteType = "Unknown" // 未知类型

	HtmlByteType       ByteType = "Html"       // HTML 类型
	JsonByteType       ByteType = "Json"       // JSON 类型
	JavaScriptByteType ByteType = "JavaScript" // JavaScript 类型

	Base64ByteType ByteType = "Base64" // Base64 类型

	FormByteType ByteType = "Form" // 表单类型  GET和POST 都是表单
)

// TypeAssertions 类型断言
type TypeAssertions interface {
	IsYou(data []byte) bool // 是否是你
	GetType() ByteType      // 获取类型
}

// DataAssertions 数据断言
type DataAssertions interface {
	Assertions([]byte) ByteType
}
