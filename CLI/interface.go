package CLI

import (
	"flag"
	"fmt"
)

// CLI 接口规范
// 程序路由 、 参数

// CLI 规范
type CLI interface {
}

// XCheck -x-cli -poc cve-2023-40787 -u http://www.baidu.com -c "proxy http://127.0.0.1:8080"

// XCheck -x-server -l :9000

func Test() {

	sc := flag.String("poc", "", "poc")

	m := flag.Bool("x", false, "x")

	h := flag.String("h", "", "url")

	flag.Parse()

	if *m {

		fmt.Println(*sc)
		fmt.Println(*h)
	}

}
