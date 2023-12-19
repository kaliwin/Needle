package middleman

import (
	"crypto/tls"
	"fmt"
	"github.com/kaliwin/Needle/crypto/certificate"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// StartMiddleman 启动中间人
// 监听某个端口将该端口上的https转发给代理服务器
// 实验性功能 陷入堵塞表示服务启动成功
func StartMiddleman(serverAddr string, proxyAdder string, CACertPath string, CAKeyPath string) error {

	ca, err2 := os.ReadFile(CACertPath)
	readFile, err2 := os.ReadFile(CAKeyPath)

	if err2 != nil {
		return err2
	}
	pemCA, err2 := certificate.LoadCA(ca, readFile) // 加载CA
	if err2 != nil {
		return err2
	}

	parse, _ := url.Parse(proxyAdder)
	c := http.Client{}
	c.Transport = &http.Transport{ //
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Proxy: http.ProxyURL(parse),
	}

	server := &http.Server{
		Addr: serverAddr, // 监听端口
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.URL.Scheme = "https" // 协议头写死https
			r.URL.Host = r.Host
			r.RequestURI = ""
			//r.Host = "www.baidu.com"

			do, err := c.Do(r)
			if err != nil {
				log.Println(err)
				return
			}

			for k, v := range do.Header { // 设置头
				if v != nil {
					w.Header().Set(k, v[0])
				}
			}

			all, _ := io.ReadAll(do.Body)
			w.WriteHeader(do.StatusCode) // 设置状态码
			w.Write(all)                 // 写回响应体
		}),
		TLSConfig: &tls.Config{
			GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
				fmt.Println(info.ServerName + " RemoteAddr: " + info.Conn.RemoteAddr().String())
				//pair, err := tls.LoadX509KeyPair("/root/cyvk/ManDown/go-tmp/static/cert.pem", "/root/cyvk/ManDown/go-tmp/static/key.pem")
				// 签发证书
				cert, privacy, err := certificate.MakePemCert(pemCA, []string{"*." + info.ServerName, info.ServerName}, info.ServerName, certificate.GetTlsCASubject())
				if err != nil {
					log.Println(err)
					return nil, err
				}
				pair, err := tls.X509KeyPair(cert, privacy)
				if err != nil {
					log.Println(err)
					return nil, err
				}
				return &pair, nil
			},
		},
	}

	fmt.Println("Server is running on " + serverAddr)
	err := server.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

	return nil
}

//// StartMiddlemanDemo 中间人演示
//func StartMiddlemanDemo()  {
//
//}
