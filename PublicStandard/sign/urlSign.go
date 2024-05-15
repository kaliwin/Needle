package sign

import (
	"crypto/md5"
	"fmt"
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
		sign += fmt.Sprintf("_%d", StrToNum(s))
	}

	return sign
}

// StrToNum 字符串转数字
// 将字符串转为字节流，然后取md5的后2位，将其加到字节流后面，然后将字节流的每个字节相加
// 确保每个字符串的签名都是唯一的
func StrToNum(str string) int {
	bytes := []byte(str)
	sum := md5.Sum([]byte(str))
	bytes = append(bytes, sum[14:]...)
	t := 0
	for _, b := range bytes {
		t += int(b)
	}
	return t
}
