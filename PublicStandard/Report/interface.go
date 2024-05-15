package Report

// 报告标准

// SendReport 发送报告
type SendReport interface {
	// SendReport 发送报告
	SendReport(url string, report ReportX) error
}

// 报告要能序列化为字节流
