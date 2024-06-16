package FieldAggregation

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"net/url"
	"os"
	"strings"
)

// ResFieldAggregation 响应字段聚合
type ResFieldAggregation struct {
	path string
	m    map[string]resMd5
}

// Accepting 接受数据 错误不会影响先前数据和后续数据
func (r ResFieldAggregation) Accepting(reqAndRes *HttpStructureStandard.HttpReqAndRes) error {

	bytes := reqAndRes.GetRes().GetData()[reqAndRes.GetRes().GetBodyIndex():]
	if len(bytes) > 1 {
		uri := reqAndRes.GetReq().GetUrl()
		parse, err := url.Parse(uri)
		if err != nil {
			return err
		}

		sum := md5.Sum(bytes)
		toString := hex.EncodeToString(sum[:])

		host := fmt.Sprintf("%s:%d", reqAndRes.GetReq().GetHttpReqService().GetIp(), reqAndRes.GetReq().GetHttpReqService().GetPort())

		r2 := r.m[toString]

		filePath := r.path + "/" + host // 文件路径

		if r2.isExist { // 响应体存在
			if r2.Host[host] { // 主机存在
				return nil
			}
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return err
			}
			r2 = resMd5{
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
		}

		fileName := fmt.Sprintf("%s-%d-%s", toString, uuid.New().ID(), path)

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

// NewResFieldAggregation 创建响应字段聚合
func NewResFieldAggregation(path string) ResFieldAggregation {
	return ResFieldAggregation{path: path, m: make(map[string]resMd5)}
}

// resMd5 响应体md5 结构体
type resMd5 struct {
	RestBodyMd5 string
	Host        map[string]bool // 主机加端口
	isExist     bool            // 存在
}
