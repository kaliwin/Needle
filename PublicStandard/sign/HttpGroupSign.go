package sign

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"github.com/kaliwin/Needle/network/http"
	"net/url"
)

// HttpGroupSign http组签名
func HttpGroupSign(list *HttpStructureStandard.HttpReqAndRes) (string, error) {

	reqSign, err := HttpReqSign(list.GetReq())
	if err != nil {
		return "", err
	}
	data := list.GetRes().GetData()
	resBodySign := "null"
	if len(data) > 1 {
		resBodySign = BodySign(data[list.GetRes().GetBodyIndex():])
	}

	return fmt.Sprintf("H-%s-B_%s-%s", reqSign, resBodySign, UuidSign()), err
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

// BodySign 体签名
func BodySign(b []byte) string {
	var d []byte
	bytes := md5.Sum(b)
	d = append(d, bytes[12:]...)
	return hex.EncodeToString(d)
}

func UuidSign() string {
	return fmt.Sprintf("T_%d", uuid.New().ID())
}
