package main

import (
	"fmt"
	"io"
	"net/http"
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

	resp, err := http.Get("https://132.120.192.66:38089/cas-server-webapp-3.3.3/")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(all))
}
