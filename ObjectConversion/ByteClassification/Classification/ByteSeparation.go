package Classification

import (
	"github.com/kaliwin/Needle/IO/Interface"
	"github.com/kaliwin/Needle/IO/file"
	"github.com/kaliwin/Needle/ObjectConversion/ByteClassification"
	"github.com/kaliwin/Needle/ObjectConversion/ByteClassification/DataType"
	"os"
	"unicode/utf8"
)

// 字节分类

// ByteSeparation 字节分类
// 基本分类 分离字符串和非字符串
type ByteSeparation struct {
	Dir       string // 分类目录
	TarGetDir string // 目标目录
}

// Do 执行分类
func (b ByteSeparation) Do() error {
	stream, err := file.BuildFIleObjectStream(b.TarGetDir, true, Interface.ObjectTypeFileByte)
	if err != nil {
		return err
	}
	list := stream.FileList

	byteType := DataType.NewStringByteType()

	for _, s := range list { // 迭代文件列表

		readFile, err := os.ReadFile(s) // 读取文件
		if err != nil {
			return err
		}
		stat, _ := os.Stat(s)
		fileName := stat.Name() // 获取文件名

		if utf8.Valid(readFile) { // 判断是否是utf8编码

			if class := byteType.Assertions(readFile); class != ByteClassification.UnknownByteType { // 执行字符串的分类
				_ = os.MkdirAll(b.Dir+"/"+string(class), 0777)
				err := os.WriteFile(b.Dir+"/"+string(class)+"/"+fileName, readFile, 0777)
				if err != nil {
					return err
				}
				continue
			}

			err := os.WriteFile(b.Dir+"/"+string(ByteClassification.StringByteType)+"/"+fileName, readFile, 0777)
			if err != nil {
				return err
			}

		} else { // 不是utf8编码
			err := os.WriteFile(b.Dir+"/"+string(ByteClassification.UnknownByteType)+"/"+fileName, readFile, 0777)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// NewByteClassification 创建字节分类
func NewByteClassification(Dir, TarGetDir string) (ByteSeparation, error) {
	err := os.MkdirAll(Dir+"/"+string(ByteClassification.StringByteType), 0777) // 创建字符串目录
	if err != nil {
		return ByteSeparation{}, err
	}
	err = os.MkdirAll(Dir+"/"+string(ByteClassification.UnknownByteType), 0777) // 创建未知目录

	if err != nil {
		return ByteSeparation{}, err
	}

	return ByteSeparation{Dir: Dir, TarGetDir: TarGetDir}, nil
}

//func StringClassify(data []byte) ByteType {
//
//}
