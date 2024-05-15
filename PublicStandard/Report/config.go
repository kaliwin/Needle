package Report

// ReportX 单个报告
type ReportX struct {
	ID          string      // 架构内唯一标识符
	Author      string      // 作者
	Date        string      // 创建日期
	Name        string      // 名称
	Description string      // 描述
	Label       []string    // 标签
	Reference   []string    // 参考外链
	Relevance   []string    // 关联程序id
	ContentType ContentType // 内容类型
	Content     any         // 内容
}

// ContentType 内容类型
type ContentType string
