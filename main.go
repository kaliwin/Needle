package main

import (
	"fmt"
	"github.com/kaliwin/Needle/PublicStandard/sign"
	"github.com/kaliwin/Needle/io/http"
	"time"
)

func main() {

	d := time.Now()
	group, err := http.ReadHttpGroupList("/root/tmp/jj/4.proto")
	if err != nil {
		panic(err)
	}

	for _, res := range group.GetHttpRawByteStreamList() {
		idSign, _ := sign.HttpBleveIdSign(res)
		fmt.Println(idSign)
		fmt.Println(len(idSign))
	}

	s := time.Now()
	fmt.Println(s.Sub(d))

}
