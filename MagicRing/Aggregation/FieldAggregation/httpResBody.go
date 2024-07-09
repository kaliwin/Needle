package FieldAggregation

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"github.com/kaliwin/Needle/PublicStandard/sign/HashingSign"
	"net/url"
	"os"
	"strings"
)

// ResFieldAggregation 响应字段聚合
// 文件名是sha256-最后一个/后的内容
type ResFieldAggregation struct {
	path string
	m    map[string]resSha256
}

// Accepting 接受数据 错误不会影响先前数据和后续数据
func (r ResFieldAggregation) Accepting(reqAndRes *HttpStructureStandard.HttpReqAndRes) error {
	uri := reqAndRes.GetReq().GetUrl()
	bytes := reqAndRes.GetRes().GetData()[reqAndRes.GetRes().GetBodyIndex():]
	if len(bytes) > 1 {

		parse, err := url.Parse(uri)
		if err != nil {
			return err
		}
		//

		sum := sha256.Sum256(bytes)
		toString := hex.EncodeToString(sum[:])

		host := fmt.Sprintf("%s:%d", reqAndRes.GetReq().GetHttpReqService().GetIp(), reqAndRes.GetReq().GetHttpReqService().GetPort())

		r2 := r.m[toString]

		filePath := r.path + "/" + host // 文件路径

		if r2.isExist { // 响应体存在
			if r2.Host[host] { // 主机存在
				return nil
			}
		} else {
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return err
			}
			r2 = resSha256{
				RestBodyMd5: toString,
				Host:        make(map[string]bool),
				isExist:     true,
			}
		}

		lastIndex := strings.LastIndex(parse.Path, "/")

		var path string

		if lastIndex == -1 && len(parse.Path) == lastIndex+1 {
			path = "index"
		} else {
			path = parse.Path[lastIndex+1:] // 取最后一个/后的内容

			if len(path) > 15 { //长度大于15 就不要了
				path = "long"
			}
		}

		fileName := fmt.Sprintf("%s_%s-%s", HashingSign.HttpResBodySha256, toString, path)

		err = os.WriteFile(filePath+"/"+fileName, bytes, os.ModePerm)
		if err != nil {
			return err
		}

		r2.Host[host] = true
		r.m[toString] = r2
		return nil

	}
	return nil
}

// NewResBodyFieldAggregation 创建响应体字段聚合
// 接受到数据会直接写入文件
func NewResBodyFieldAggregation(path string) ResFieldAggregation {
	return ResFieldAggregation{path: path, m: make(map[string]resSha256)}
}

// resSha256 响应体md5 结构体
type resSha256 struct {
	RestBodyMd5 string
	Host        map[string]bool // 主机加端口
	isExist     bool            // 存在
}
