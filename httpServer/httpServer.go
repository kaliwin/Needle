package main

import "github.com/kaliwin/Needle/httpServer/middleman"

func main() {
	//
	err := middleman.StartMiddleman(":8081", "http://127.0.0.1:8080", "/root/cyvk/ManDown/CA/burpCA.cer", "/root/cyvk/ManDown/CA/burpCA-key.cer")
	if err != nil {
		panic(err)
	}

	//go middleman.StartMiddleman(":443", "http://127.0.0.1:8080", "/root/cyvk/ManDown/CA/burpCA.cer", "/root/cyvk/ManDown/CA/burpCA-key.cer")

	//cac, _ := os.ReadFile("/root/tmp/burpca")
	//ketc, _ := os.ReadFile("/root/tmp/burpkey")
	//
	//pemCA, err2 := certificate.LoadCA(cac, ketc) // 加载CA
	//if err2 != nil {
	//	panic(err2)
	//}
	//
	//certificate.MakePemCert(pemCA, []string{"www.baidu.com"}, "sd", certificate.GetTlsCASubject())

}
