package main

import "github.com/kaliwin/Needle/httpServer/middleman"

func main() {

	err := middleman.StartMiddleman(":443", "http://127.0.0.1:8080", "/root/tmp/burpCA.cer", "/root/tmp/burpCA-key.cer")
	if err != nil {
		panic(err)
	}

}
