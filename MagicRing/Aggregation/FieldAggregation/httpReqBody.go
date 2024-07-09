package FieldAggregation

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"github.com/kaliwin/Needle/PublicStandard/sign/HashingSign"
	"github.com/kaliwin/Needle/network/http"
	"net/url"
	"os"
	"strings"
)

// http组 精简聚合

// http 请求组聚合 体和响应体是一样的 在于提交表单的额外分析

type ReqBodyFieldAggregation struct {
	httpFormAggregate httpFormAggregate // 表单聚合
	path              string            // 路径
	isExist           map[string]bool   // 是否存在
}

// Accepting 接受数据 错误不会影响先前数据和后续数据
// 只会处理表单数据 多表单暂时搁置
func (r *ReqBodyFieldAggregation) Accepting(Accepting *HttpStructureStandard.HttpReqAndRes) error {

	req := Accepting.GetReq()

	Host := fmt.Sprintf("%s:%d", req.GetHttpReqService().GetIp(), req.GetHttpReqService().GetPort())

	httpReq, err := http.BuildRefactorStandardHttpReq(req, nil)
	if err != nil {
		return err
	}

	request := httpReq.BuildRequest()

	ContentType := request.Header.Get("Content-Type") // 获取类型 这是一种表单标准

	query := request.URL.Query()

	uriC, _ := url.Parse(Accepting.GetReq().GetUrl()) // 解析URI

	uri := fmt.Sprintf("%s:%d/%s", Accepting.GetReq().GetHttpReqService().GetIp(), Accepting.GetReq().GetHttpReqService().GetPort(), uriC.Path) // 获取URI

	for k, v := range query {
		r.httpFormAggregate.httpForm[uri] = append(r.httpFormAggregate.httpForm[uri], HttpForm{k: v})
	}

	if strings.Contains(ContentType, "application/x-www-form-urlencoded") { // 表单
		_ = request.ParseForm()
		for k, v := range request.Form {
			r.httpFormAggregate.httpForm[uri] = append(r.httpFormAggregate.httpForm[uri], HttpForm{k: v})
		}
	}
	// 分段表单 暂时搁置不处理
	//if strings.Contains(ContentType, "multipart/form-data") { // 分段表单 通常用于文件上传
	//	//_ = request.ParseMultipartForm(1024 * 1024 * 30)
	//	//fmt.Println(request.MultipartForm.Value)
	//}

	data := req.GetData()[req.GetBodyIndex():]

	if len(data) > 0 { // 存在请求体
		sum256 := sha256.Sum256(data)
		reqSha256 := hex.EncodeToString(sum256[:])

		index := strings.LastIndex(httpReq.GetPath(), "/")
		getPath := httpReq.GetPath()
		name := ""

		if index+1 == len(getPath) {
			name = "index"
		} else {
			name = getPath[index+1:]
			if len(name) > 15 { //长度大于15 就不要了
				name = "long"
			}

		}

		//getPath := httpReq.GetPath()[strings.LastIndex(httpReq.GetPath(), "/"):]
		fileName := Host + "/" + HashingSign.HttpReqBodySha256 + "_" + reqSha256 + "_" + name

		if !r.isExist[fileName] {
			_ = os.MkdirAll(r.path+"/"+Host, os.ModePerm)
			err := os.WriteFile(r.path+"/"+fileName, data, os.ModePerm)
			if err != nil {
				return err
			}
			r.isExist[fileName] = true
		}
	}

	return nil
}

// HttpForm 表单
type HttpForm map[string][]string

// NewReqBodyFieldAggregation 创建请求体聚合
func NewReqBodyFieldAggregation(path string) ReqBodyFieldAggregation {
	return ReqBodyFieldAggregation{
		httpFormAggregate: httpFormAggregate{
			httpForm: make(map[string][]HttpForm),
		},
		isExist: make(map[string]bool),
		path:    path,
	}
}

type httpFormAggregate struct {
	httpForm map[string][]HttpForm // Key 是URI Value 是表单列表
}
