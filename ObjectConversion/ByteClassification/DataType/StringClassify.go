package DataType

import "github.com/kaliwin/Needle/ObjectConversion/ByteClassification"

// StringByteType 字符串类型分类
type StringByteType struct {
	typeAssertions []ByteClassification.TypeAssertions
}

// Assertions 断言 迭代所有的断言 如果有一个断言为真 则返回该类型 否则返回未知类型
func (s StringByteType) Assertions(bytes []byte) ByteClassification.ByteType {
	for _, assertion := range s.typeAssertions {
		if assertion.IsYou(bytes) {
			return assertion.GetType()
		}
	}
	return ByteClassification.UnknownByteType
}

func NewStringByteType() ByteClassification.DataAssertions {
	byteType := StringByteType{}
	byteType.typeAssertions = append(byteType.typeAssertions, NewHtmlTypeAssertions())       // HTML 类型
	byteType.typeAssertions = append(byteType.typeAssertions, NewJavaScriptTypeAssertions()) // JavaScript 类型
	byteType.typeAssertions = append(byteType.typeAssertions, NewJsonTypeAssertions())       // JSON 类型
	byteType.typeAssertions = append(byteType.typeAssertions, NewBaseTypeAssertions())       // Base 类型

	return byteType
}
