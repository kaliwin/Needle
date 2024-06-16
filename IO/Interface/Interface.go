package Interface

// IO 标准接口

// Iteration 迭代器
type Iteration interface {
	Next() (any, error)
	Class() ObjectType
	Length() int // 数据长度
	Close() error
	Iteration(func(any) bool)
}
