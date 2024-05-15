package ObjectHandling

import (
	"errors"
	Interface "github.com/kaliwin/Needle/MagicRing/Integrate"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"google.golang.org/protobuf/proto"
	"io"
	"os"
)

// ObjectStreamRead 对象流读取的具体实现

// FileIOReadStream 文件IO读取流
type FileIOReadStream struct {
	ObjectType Interface.ObjectType // 对象类型
	FilePath   string               // 文件路径
	IsDir      bool                 // 是否是目录
	FileList   []string             // 文件列表
	subscript  int                  // 下标
}

func (f *FileIOReadStream) Next() (any, error) {
	if f.subscript+1 > len(f.FileList) {
		return nil, io.EOF
	}
	s := f.FileList[f.subscript]
	f.subscript++
	file, err := os.ReadFile(s)
	if err != nil {
		return nil, err
	}

	return Conversion(file, f.ObjectType) // 调用Conversion函数 进行数据转换
}

func (f *FileIOReadStream) Length() int {
	return len(f.FileList)
}

// Close 关闭流
func (f *FileIOReadStream) Close() error {
	f.FileList = nil
	f.subscript = 1
	return nil
}

// Go 启动 如果是目录则读取目录下的所有文件
func (f *FileIOReadStream) Go() error {
	//f.subscript = -1
	if f.IsDir { // true 读取目录
		ts := make([]string, 0)
		err := ShowAllFile(f.FilePath, &ts)
		if err != nil {
			return err
		}
		f.FileList = ts
	} else { // false 读取文件
		f.FileList = append(f.FileList, f.FilePath)
	}
	return nil
}

// ShowAllFile 递归读取目录下的所有文件
func ShowAllFile(dir string, res *[]string) error {
	readDir, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range readDir {
		if entry.IsDir() { // 如果是目录 递归
			return ShowAllFile(dir+"/"+entry.Name(), res)
		} else {
			*res = append(*res, dir+"/"+entry.Name())
		}
	}
	return nil
}

// 数据转换

// Conversion 数据转换 标记器仅支持[]byte 和 protobuf反序列化后的实例 两种数据类型
func Conversion(data []byte, ObjectType Interface.ObjectType) (any, error) {
	switch ObjectType {
	case Interface.ObjectTypeFileByte: // 文件字节流
		return data, nil

	case Interface.ObjectTypeHttpGroupProto: // http组协议
		httpGroup := HttpStructureStandard.HttpReqAndRes{}
		err := proto.Unmarshal(data, &httpGroup)
		return &httpGroup, err
	case Interface.ObjectTypeHttpGroupListProto: // http原始字节流列表
		list := HttpStructureStandard.HttpRawByteStreamList{}
		err := proto.Unmarshal(data, &list)
		return &list, err
	}

	return nil, errors.New("unsupported data type")
}

// BuildFIleObjectStream 构建文件对象流
func BuildFIleObjectStream(path string, isDir bool, objectType Interface.ObjectType) (FileIOReadStream, error) {
	stream := FileIOReadStream{
		ObjectType: objectType,
		FilePath:   path,
		IsDir:      isDir,
	}
	return stream, stream.Go()
}
