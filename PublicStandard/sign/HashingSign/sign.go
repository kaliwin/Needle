package HashingSign

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
)

// 散列码类型

// 开头一个字符大写

const Separator = "_"

const (
	HttpGroupSha256   = "HG" // 组散列
	HttpReqDataSha256 = "QD"
	HttpResDataSha256 = "SD"
	HttpReqBodySha256 = "QB" // 请求体散列
	HttpResBodySha256 = "SB" // 响应体散列
)

// Sha256HttpGroup 散列http组
// 请求报文加上响应报文加上目标地址 求出散列值
func Sha256HttpGroup(data *HttpStructureStandard.HttpReqAndRes) string {

	bytes := append(data.GetReq().GetData(), data.GetRes().GetData()...)

	host := fmt.Sprintf("%s:%d", data.GetReq().GetHttpReqService().GetIp(), data.GetReq().GetHttpReqService().GetPort())

	bytes = append(bytes, []byte(host)...)
	sum256 := sha256.Sum256(bytes)

	return HttpGroupSha256 + Separator + hex.EncodeToString(sum256[:])
}

// Sha256HttpReq 散列http请求
func Sha256HttpReq(data *HttpStructureStandard.HttpReqData) string {
	sum256 := sha256.Sum256(data.GetData())
	return HttpReqDataSha256 + Separator + hex.EncodeToString(sum256[:])
}

// Sha256HttpRes 散列http响应
func Sha256HttpRes(data *HttpStructureStandard.HttpResData) string {
	sum256 := sha256.Sum256(data.GetData())
	return HttpResDataSha256 + Separator + hex.EncodeToString(sum256[:])
}

// Sha256HttpReqBody 散列请求体
func Sha256HttpReqBody(data *HttpStructureStandard.HttpReqData) string {
	sum256 := sha256.Sum256(data.GetData()[data.GetBodyIndex():])
	return HttpReqBodySha256 + Separator + hex.EncodeToString(sum256[:])
}

// Sha256HttpResBody 散列请求体
func Sha256HttpResBody(data *HttpStructureStandard.HttpResData) string {
	sum256 := sha256.Sum256(data.GetData()[data.GetBodyIndex():])
	return HttpResBodySha256 + Separator + hex.EncodeToString(sum256[:])
}
