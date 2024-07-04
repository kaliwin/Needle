package middleman

import (
	"crypto/tls"
	"github.com/kaliwin/Needle/crypto/certificate"
	"net/http"
)

type HttpsTest struct {
	//Server http.Server
	CACert certificate.CACert // CA证书 用于动态签发通信证书 通常burp可以设置信任所有证书
}

func (receiver *HttpsTest) Go() {

	server := http.Server{}

	server.Addr = ":8081"
	server.TLSConfig = &tls.Config{
		GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
			return BuildCa(receiver.CACert, info.ServerName)
		},
	}

	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(r.URL.Path))

	})

	//server.Serve()

	server.ListenAndServeTLS("", "")

}
