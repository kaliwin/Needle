package main

import (
	"fmt"
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

	//ca, i, err := certificate.MakeCA(certificate.Rsa2048, certificate.GetTlsCASubject())
	//if err != nil {
	//	panic(err)

	//cac, _ := os.ReadFile("/root/tmp/CA")
	//i, _ := os.ReadFile("/root/tmp/CAKey")
	//
	//loadCA, err := certificate.LoadPemCA(cac, i)
	//if err != nil {
	//	panic(err)
	//}
	//cert, privacy, err := certificate.MakePemCert(loadCA, []string{"erin.server"}, "erin.server", certificate.GetTlsCASubject())
	//if err != nil {
	//	panic(err)
	//}
	//os.WriteFile("/root/tmp/client.crt", cert, os.ModePerm)
	//os.WriteFile("/root/tmp/client.key", privacy, os.ModePerm)
	//sc, _ := os.ReadFile("/root/tmp/client.crt")
	//sk, _ := os.ReadFile("/root/tmp/client.key")
	//
	//serverCert, err := tls.X509KeyPair(sc, sk)
	//
	////创建 tls.Certificate 对象
	//tlsCert := tls.Certificate{
	//	Certificate: serverCert.Certificate,
	//	PrivateKey:  serverCert.PrivateKey,
	//}

	//data, err := pkcs12.Modern.Encode()
	//if err != nil {
	//	panic(err)
	//}
	//os.WriteFile("/root/tmp/CA.p12", data, os.ModePerm)

	fmt.Println("hello world!")
}
