package IO

import (
	"errors"
	"github.com/kaliwin/Needle/IO/Interface"
	"github.com/kaliwin/Needle/IO/file"
)

// BuildResourceDescriptionRead 构建资源描述读取流
func BuildResourceDescriptionRead(uri Interface.ResourceDescription) (Interface.Iteration, error) {
	switch uri.Protocol {
	case Interface.IOFile:
		stream, err := file.BuildFIleObjectStream(uri.Path, true, uri.ObjectType)
		return &stream, err
	}
	return nil, errors.New("not support protocol")
}
