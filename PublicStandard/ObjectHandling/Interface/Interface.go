package Interface

// 接口规范

// ObjectStreamRead 对象流读取
type ObjectStreamRead interface {
	Next() (any, error) // 读取数据
	Length() int        // 数据长度
	Close() error       // 关闭流
}
