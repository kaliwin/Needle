package DataType

import (
	"github.com/dop251/goja"
	"github.com/kaliwin/Needle/ObjectConversion/ByteClassification"
)

// JavaScriptTypeAssertions JavaScript类型断言
// 语法必须要标准 每行要用; 结尾 [?]发现goja无法解析 但是流量器可以解析 有待研究
type JavaScriptTypeAssertions struct {
}

func (j JavaScriptTypeAssertions) IsYou(data []byte) bool {
	_, err := goja.Parse("test.js", string(data))
	if err != nil {
		return false
	}
	return true
}

func (j JavaScriptTypeAssertions) GetType() ByteClassification.ByteType {
	return ByteClassification.JavaScriptByteType
}

func NewJavaScriptTypeAssertions() ByteClassification.TypeAssertions {
	return JavaScriptTypeAssertions{}
}
