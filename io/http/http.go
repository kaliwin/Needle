package http

import (
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"google.golang.org/protobuf/proto"
	"os"
)

// ReadHttpGroup 读取http组
func ReadHttpGroup(path string) (*HttpStructureStandard.HttpReqAndRes, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	httpGroup := HttpStructureStandard.HttpReqAndRes{}
	return &httpGroup, proto.Unmarshal(file, &httpGroup)
}

// ReadHttpGroupList 读取http组列表
func ReadHttpGroupList(path string) (*HttpStructureStandard.HttpRawByteStreamList, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	httpGroupList := HttpStructureStandard.HttpRawByteStreamList{}
	return &httpGroupList, proto.Unmarshal(file, &httpGroupList)
}
