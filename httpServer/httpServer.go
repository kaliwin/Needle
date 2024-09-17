package main

import (
	"github.com/kaliwin/Needle/httpServer/middleman"
	"github.com/kaliwin/Needle/network/dns"
	"net/http"
	"time"
)

func Test() {
	//burpCA, _ := os.ReadFile("/root/cyvk/ManDown/CA/burpCA.cer")
	//burpCAKey, _ := os.ReadFile("/root/cyvk/ManDown/CA/burpCA-key.cer")

	acting, err := middleman.BuildActing("http://127.0.0.1:8080", "/root/cyvk/ManDown/CA/burpCA.cer", "/root/cyvk/ManDown/CA/burpCA-key.cer")
	if err != nil {
		panic(err)
	}

	yi := middleman.Yi{
		Lock: make(chan bool, 1),
	}
	acting.Client = &yi

	server := http.Server{
		Addr:    ":8081",
		Handler: acting,
	}
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second * 1)

	yi.Go()

	time.Sleep(time.Hour * 1)
}

func main() {
	////
	//err := middleman.StartMiddleman(":8081", "http://127.0.0.1:8080", "/root/cyvk/ManDown/CA/burpCA.cer", "/root/cyvk/ManDown/CA/burpCA-key.cer")
	//if err != nil {
	//	panic(err)
	//}

	//go middleman.StartMiddleman(":443", "http://127.0.0.1:8080", "/root/cyvk/ManDown/CA/burpCA.cer", "/root/cyvk/ManDown/CA/burpCA-key.cer")
	//
	////cac, _ := os.ReadFile("/root/tmp/burpca")
	////ketc, _ := os.ReadFile("/root/tmp/burpkey")
	////
	////pemCA, err2 := certificate.LoadCA(cac, ketc) // 加载CA
	////if err2 != nil {
	////	panic(err2)
	////}
	////
	////certificate.MakePemCert(pemCA, []string{"www.baidu.com"}, "sd", certificate.GetTlsCASubject())
	////Test()
	////middleman.BuildObjectC()
	////dns.ServeDNS(":53", "puniaokeji.com", "192.168.124.11")
	//
	//httpProxyImport, err := proxyImport.BuildHttpProxyImport("/root/tmp/nginx", "/root/cyvk/ManDown/CA/burpCA.cer", "/root/cyvk/ManDown/CA/burpCA-key.cer")
	//if err != nil {
	//	panic(err)
	//}
	//
	//now := time.Now()
	//
	//go httpProxyImport.StartHttpServer(":9010")
	////time.Sleep(time.Second * 1)
	//httpProxyImport.Go("http://127.0.0.1:8080")
	//now1 := time.Now()
	//fmt.Println(now1.Sub(now).String())
	//fmt.Println("进程结束")
	//time.Sleep(time.Hour * 8)
	//fmt.Println("进程结束")

	//ca, err := os.ReadFile("/root/cyvk/ManDown/CA/burpCA.cer")
	//if err != nil {
	//	panic(err)
	//}
	//
	//key, err := os.ReadFile("/root/cyvk/ManDown/CA/burpCA-key.cer")
	//if err != nil {
	//	panic(err)
	//}
	//
	//loadCA, _ := certificate.LoadCA(ca, key)
	//test := middleman.HttpsTest{CACert: loadCA}
	//go test.Go()

	//
	err := dns.ServeDNS(":53", []string{"xcheck"}, "192.168.3.108")

	//err := middleman.StartMiddleman(":443", "http://127.0.0.1:12333", "/root/cyvk/ManDown/CA/burpCA.cer", "/root/cyvk/ManDown/CA/burpCA-key.cer")
	if err != nil {
		panic(err)
	}

	//fmt.Println("dd")

	//middlemanHttp.HttpServer.Addr = ":9010"
	//middlemanHttp.HttpServer.ListenAndServe()
}
