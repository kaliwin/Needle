package sign

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

// UrlSign	Url签名
// 第一位为Host的签名，后面为Path的签名最大4个字符 一共21位
func UrlSign(Url *url.URL) string {
	sign := fmt.Sprintf("%s", WuSign([]byte(Url.Host)))
	//fmt.Println(sign)
	size := 4
	i := 0
	var d = []string{"0000", "0000", "0000", "0000"}

	for _, s := range strings.Split(Url.Path, "/") {
		if s == "" {
			continue
		}
		if i+1 > size {
			break
		}

		//sign += WuSign(s)
		d[i] = PathSign([]byte(s))
		i++
	}
	join := strings.Join(d, "")
	return sign + join
}

// WuSign	Url Path签名
// md5 计算 得到 5位字符的签名
func WuSign(s []byte) string {
	sum := md5.Sum(s)
	t := hex.EncodeToString(sum[2:3])        // 取md5的第三位 hex编码 前两位
	toString := hex.EncodeToString(sum[15:]) // 取md5的在最后一位 hex编码 后两位
	e := hex.EncodeToString(sum[1:2])[:1]    // 取md5的第二位 hex编码后的第一个字符 最后一位

	return t + toString + e

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

// PathSign	路径签名 4个字符
func PathSign(str []byte) string {
	sum := md5.Sum(str)
	toString := hex.EncodeToString(sum[15:]) // 取md5的在最后一位 hex编码
	t := hex.EncodeToString(sum[2:3])        // 取md5的第三位 hex编码
	return t + toString
}
