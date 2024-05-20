package main

import (
	"fmt"
	"github.com/kaliwin/Needle/PublicStandard/sign"
	"net/url"
)

func main() {

	//d := time.Now()
	//group, err := http.ReadHttpGroupList("/root/tmp/jj/4.proto")
	//if err != nil {
	//	panic(err)
	//}

	//for _, res := range group.GetHttpRawByteStreamList() {
	//	idSign, _ := sign.HttpBleveIdSign(res)
	//	fmt.Println(idSign)
	//	fmt.Println(len(idSign))
	//}

	parse, _ := url.Parse("http://localhost:8080")
	fmt.Println(sign.UrlSign(parse))
	fmt.Println(len(sign.UrlSign(parse)))
	fmt.Println(sign.ThirteenSign([]byte("http://localhost:8080")))
	fmt.Println(len(sign.ThirteenSign([]byte("http://localhost:8080"))))
}
