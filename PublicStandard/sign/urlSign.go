package sign

import (
	"fmt"
	"github.com/google/uuid"
	"net/url"
	"strings"
)

// UrlSign	Url签名
// 第一位为Host的签名，后面为Path的签名，最后为uuid的ID
func UrlSign(Url *url.URL) string {
	sign := fmt.Sprintf("%d", StrToNum(Url.Host))
	for _, s := range strings.Split(Url.Path, "/") {
		if s == "" {
			continue
		}
		sign += fmt.Sprintf("-%d", StrToNum(s))
	}

	return fmt.Sprintf("%s-%d", sign, uuid.New().ID())
}

// StrToNum 字符串转数字
// 将字符串转为字节流，然后将字节流的每个字节相加，最后除以9
func StrToNum(str string) int {
	bytes := []byte(str)
	t := 0
	for _, b := range bytes {
		t += int(b)
	}
	return t / 9
}
