package sign

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"github.com/kaliwin/Needle/network/http"
	"net/url"
	"strconv"
)

// HttpBleveIdSign http bleve id签名
// 限制在 36个字符
func HttpBleveIdSign(list *HttpStructureStandard.HttpReqAndRes) (string, error) {
	getUrl := list.GetReq().GetUrl()
	parse, err := url.Parse(getUrl)
	if err != nil {
		return "", err
	}
	uriSign := UrlSign(parse)
	wuSign := WuSign(list.GetRes().GetData())
	uuidSign := UuidSign()
	return uriSign + wuSign + uuidSign, nil
}

func HttpReqSign(req *HttpStructureStandard.HttpReqData) (string, error) {
	httpReq, err := http.BuildRefactorStandardHttpReq(req, nil)
	if err != nil {
		return "", err
	}
	method := httpReq.GetMethod()[:1]
	parse, _ := url.Parse(req.GetUrl())
	sign := UrlSign(parse)
	return fmt.Sprintf("R_%s_%s", method, sign), nil
}

// UuidSign	唯一签名 10位字符
func UuidSign() string {
	itoa := strconv.Itoa(int(uuid.New().ID()))
	c := len(itoa)
	for i := 0; i < 10-c; i++ {
		itoa += strconv.Itoa(i)
	}
	return itoa
}
