package sign

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"github.com/kaliwin/Needle/network/http"
	"net/url"
	"strconv"
)

// HttpBleveIdSign http bleve id签名
// 限制在 36个字符
//func HttpBleveIdSign(list *HttpStructureStandard.HttpReqAndRes) (string, error) {
//
//	getUrl := fmt.Sprintf("%s:%d", list.GetReq().GetHttpReqService().GetIp(), list.GetReq().GetHttpReqService().GetPort())
//	uriSign := ThirteenSign([]byte(getUrl))
//	wuSign := ThirteenSign(list.GetRes().GetData())
//	uuidSign := UuidSign()
//	return uriSign + wuSign + uuidSign, nil
//}

// HttpBleveIdSign 6.27 更新 使用sha256签名
func HttpBleveIdSign(list *HttpStructureStandard.HttpReqAndRes) string {
	reqB := list.GetReq().GetData()
	resB := list.GetRes().GetData()
	tarGet := fmt.Sprintf("%s:%d-%t", list.GetReq().GetHttpReqService().GetIp(), list.GetReq().GetHttpReqService().GetPort(), list.GetReq().GetHttpReqService().GetSecure())
	reqB = append(reqB, resB...)
	reqB = append(reqB, []byte(tarGet)...)
	sum256 := sha256.Sum256(reqB)
	return hex.EncodeToString(sum256[:])
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
