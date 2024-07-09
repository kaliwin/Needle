package DataType

import (
	"bytes"
	"github.com/kaliwin/Needle/ObjectConversion/ByteClassification"
	"golang.org/x/net/html"
)

// HtmlTypeAssertions html类型断言
type HtmlTypeAssertions struct {
}

// IsYou 使用html语法解析器 如果开头不是 <.*> 的话 返回false
func (h HtmlTypeAssertions) IsYou(data []byte) bool {

	p := html.NewTokenizer(bytes.NewReader(data))

	if t := p.Next(); t == html.ErrorToken || t == html.TextToken {
		return false
	}

	return true
}

// GetType 返回类型
func (h HtmlTypeAssertions) GetType() ByteClassification.ByteType {
	return ByteClassification.HtmlByteType
}

func NewHtmlTypeAssertions() ByteClassification.TypeAssertions {
	return HtmlTypeAssertions{}
}
