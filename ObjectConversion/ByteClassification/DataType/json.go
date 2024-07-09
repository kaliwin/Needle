package DataType

import (
	"encoding/json"
	"github.com/kaliwin/Needle/ObjectConversion/ByteClassification"
)

type JsonTypeAssertions struct {
}

func (j JsonTypeAssertions) IsYou(data []byte) bool {

	d := map[string]any{}

	err := json.Unmarshal(data, &d)
	if err != nil {
		return false
	}

	return true
}

// GetType 返回类型
func (j JsonTypeAssertions) GetType() ByteClassification.ByteType {
	return ByteClassification.JsonByteType
}

func NewJsonTypeAssertions() ByteClassification.TypeAssertions {
	return JsonTypeAssertions{}
}
