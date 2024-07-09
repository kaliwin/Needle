package DataType

import (
	"encoding/base64"
	"github.com/kaliwin/Needle/ObjectConversion/ByteClassification"
)

type BaseTypeAssertions struct {
}

func (b BaseTypeAssertions) IsYou(data []byte) bool {

	_, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return false
	}
	return true
}

func (b BaseTypeAssertions) GetType() ByteClassification.ByteType {
	return ByteClassification.Base64ByteType
}

func NewBaseTypeAssertions() ByteClassification.TypeAssertions {
	return BaseTypeAssertions{}
}
