package Interface

import "net"

// SearchMode 搜查模式
type SearchMode string

// DataSourcesType 数据源类型
type DataSourcesType string

// ProjectIndexType 项目索引类型
type ProjectIndexType string

// FieldValue 字段值
type FieldValue struct {
	Id    string    // id
	Type  FieldType // 类型
	Value []any     // 值
}

// FieldType 字段类型
type FieldType string

//  默认的空值

var (
	StringNull = "%StringNull%"
	IPNull     = net.ParseIP("0.0.0.1")
)

// IOType IO流类型 bleve索引库的不同IO流
// 设计之初考虑到可能存在远程IO流的需求比如存储桶流 或者FTP流等
// [?] 倾向于使用本地文件流
type IOType string // IO 流类型

// FieldType 字段类型
const (
	StringFieldType   FieldType = "string"   // 字符串类型
	NumberFieldType   FieldType = "number"   // 数字类型
	BoolFieldType     FieldType = "bool"     // 布尔类型
	DateTimeFieldType FieldType = "datetime" // 日期时间类型
	IpFieldType       FieldType = "ip"       // ip类型
)

// IOType 默认情况下会读取url指定的目录下的所有文件
const (
	IOTypeFile IOType = "IOTypeFile" // 文件IO
	IOTypeOss  IOType = "IOTypeOss"  // OSS IOw
	IOTypeFTP  IOType = "IOTypeFTP"  // FTP IO

	// IOTypeHttpIo 使用httpIO的时候需要暴露目录或提供一个fileList.txt文件
	// 用于读取文件
	IOTypeHttpIo IOType = "IOTypeHttpIo" // http IO
)

// ObjectType 对象类型 用于将字节流转换为具体的对象 然后进行标记 标记器对传入的对象有类型要求
type ObjectType string

const (
	ObjectTypeHttpGroupProto     ObjectType = "ObjectTypeHttpGroupProto"     // http组Proto类型
	ObjectTypeHttpGroupListProto ObjectType = "ObjectTypeHttpGroupListProto" // http组列表Proto类型

	//ObjectTypeFileByte  文件字节类型 单个文件 例如js html json等
	ObjectTypeFileByte ObjectType = "ObjectTypeFileByte"
)

// IndexTypeID 索引类型ID
// 每个索引库都会写入一个索引类型ID 用于区分不同的索引库
const IndexTypeID string = "IndexTypeID"

const (
	ProjectHttpIndexType ProjectIndexType = "HttpIndexProject"      // http索引项目
	CallChainIndex       ProjectIndexType = "CallChainIndexProject" // 调用链索引项目
)

// DataSourcesType 数据类型

//const (
//	ProtoHttpReqAndResList DataSourcesType = "http"          // proto http请求和响应列表
//	RawImg                 DataSourcesType = "RawImg"        // 原始图片
//	RawJavaScript          DataSourcesType = "RawJavaScript" // 原始JavaScript
//)

const MarkTypeHtml DataSourcesType = "Html" // html标记类型
const MarkTypeHttp DataSourcesType = "Http"
const MarkTypeHttpList DataSourcesType = "HttpList"
const MarkTypeJs DataSourcesType = "JavaScript"
const MarkTypeJson DataSourcesType = "JSON"
const MarkTypeQuery DataSourcesType = "QUERY"

const (
	OSS  SearchMode = "OSS"  // oss列表
	Zip  SearchMode = "Zip"  // 压缩包
	File SearchMode = "File" // 本地文件
)
