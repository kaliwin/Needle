package IO

import (
	"errors"
	"github.com/kaliwin/Needle/IO/Interface"
	"github.com/kaliwin/Needle/IO/file"
	"os"
)

// BuildResourceDescriptionRead 构建资源描述读取流
func BuildResourceDescriptionRead(uri Interface.ResourceDescription) (Interface.Iteration, error) {
	switch uri.Protocol {
	case Interface.IOFile:
		stat, err2 := os.Stat(uri.Path)
		if err2 != nil {
			return nil, err2
		}
		stream, err := file.BuildFIleObjectStream(uri.Path, stat.IsDir(), uri.ObjectType)
		return &stream, err
	}
	return nil, errors.New("not support protocol")
}
