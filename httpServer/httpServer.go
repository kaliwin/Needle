package main

import "github.com/kaliwin/Needle/httpServer/middleman"

func main() {

	go middleman.StartMiddleman(":8083", "http://127.0.0.1:8080", "/root/cyvk/ManDown/CA/burpCA.cer", "/root/cyvk/ManDown/CA/burpCA-key.cer")

	go middleman.StartMiddleman(":443", "http://127.0.0.1:8080", "/root/cyvk/ManDown/CA/burpCA.cer", "/root/cyvk/ManDown/CA/burpCA-key.cer")

	err := middleman.StartMiddleman(":8081", "http://127.0.0.1:8080", "/root/cyvk/ManDown/CA/burpCA.cer", "/root/cyvk/ManDown/CA/burpCA-key.cer")
	if err != nil {
		panic(err)
	}

}
