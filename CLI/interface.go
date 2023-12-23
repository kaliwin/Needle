package CLI

// CLI 接口规范
// 程序路由 、 参数

// Options 选项
type Options interface {
}

// XCheck -x-cli -poc cve-2023-40787 -u http://www.baidu.com -c "proxy http://127.0.0.1:8080"

// XCheck -x-server -l :9000
