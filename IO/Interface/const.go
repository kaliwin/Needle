package Interface

type ObjectType string // 对象类型 可序列化转换的对象

type IOProtocolType string // IO 协议类型

// ResourceDescription 资源描述
type ResourceDescription struct {
	Protocol   IOProtocolType    // 协议
	Path       string            // 路径
	ObjectType ObjectType        // 对象类型
	Config     map[string]string // 配置信息
}

const (
	ObjectTypeHttpGroupProto     ObjectType = "ObjectTypeHttpGroupProto"     // http组Proto类型
	ObjectTypeHttpGroupListProto ObjectType = "ObjectTypeHttpGroupListProto" // http组列表Proto类型

	//ObjectTypeFileByte  文件字节类型 单个文件 例如js html json等
	ObjectTypeFileByte ObjectType = "ObjectTypeFileByte"

	ObjectTypePcap ObjectType = "ObjectTypePcap" // pcap文件

	//
	//ObjectTypeFilePath ObjectType = "ObjectTypeFilePath" // 文件路径
)

const (
	IOFile IOProtocolType = "IOFile" // 本地文件流
	IOOss  IOProtocolType = "IOOss"  // 存储桶
	IOFtp  IOProtocolType = "IOFtp"  // ftp
	IOHttp IOProtocolType = "IOHttp" // http

)
