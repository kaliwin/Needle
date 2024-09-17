package middleman

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"github.com/kaliwin/Needle/crypto/certificate"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
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
			//fmt.Println(" ============================================== ")
			r.URL.Scheme = "https" // 协议头写死https
			r.URL.Host = r.Host
			r.RequestURI = ""
			//r.Host = "www.baidu.com"
			fmt.Println(r.URL.String())
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
			if strings.Index(r.URL.String(), "list") != -1 {
				fmt.Println(string(all))
			}

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
				//fmt.Println("签发证书成功")
				return &pair, nil
			},
			InsecureSkipVerify: true,
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

// MiddlemanHttp http中间人代理服务器
type MiddlemanHttp struct {
	CACert     certificate.CACert // CA证书 用于动态签发通信证书 通常burp可以设置信任所有证书
	HttpServer http.Server        // http服务器
}

// ServeHTTP 流量入口函数
func (m *MiddlemanHttp) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// 根据请求方法分发处理函数
	//fmt.Println("==============")
	//fmt.Println(request.Method)
	if request.Method == http.MethodConnect {
		m.handleHTTPS(writer, request)
	} else {
		m.handleHTTP(writer, request)
	}
}

// handleHTTP http的处理
func (m *MiddlemanHttp) handleHTTP(w http.ResponseWriter, r *http.Request) {
	hijacker, _ := w.(http.Hijacker)
	hijack, _, err := hijacker.Hijack() // 获取底层连接

	if err != nil {
		return
	}
	defer func() { _ = hijack.Close() }()
	//dial, err := net.Dial("tcp", "127.0.0.1:8081")
	//if err != nil {
	//	panic(err)
	//}

	//all, err := io.ReadAll(hijack)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(all))

	m.handleResponse(hijack, r)
}

// handleHTTPS https的处理
func (m *MiddlemanHttp) handleHTTPS(w http.ResponseWriter, r *http.Request) {
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	hijack, _, err := hijacker.Hijack() // 获取底层连接
	if err != nil {
		return
	}
	_, err = hijack.Write([]byte("HTTP/1.0 200 Connection established\r\n\r\n")) // 返回200 代表连接成功
	if err != nil {
		return
	}

	//  ==================================== //
	//
	//dial, err := net.Dial("tcp", "cyvk.server:8081")
	//if err != nil {
	//	panic(err)
	//}
	//
	//go io.Copy(dial, hijack)
	//io.Copy(hijack, dial)
	//  ==================================== //
	//
	server := tls.Server(hijack, &tls.Config{ // 设置tls配置 强制使用http1.1
		GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
			return BuildCa(m.CACert, info.ServerName) // 使用CA签发证书
		},
		NextProtos: []string{"http/1.1"}, // 强制使用http1.1 TLS协议保持默认
	})
	defer func() { _ = server.Close() }() // 关闭连接 因为是http1.1 所以每次请求都会关闭连接

	request, err := http.ReadRequest(bufio.NewReader(server)) // 读取请求
	if err != nil {
		log.Println(err)
		return
	}
	m.handleResponse(server, request)
}

// handleResponse 处理响应
func (m *MiddlemanHttp) handleResponse(w io.Writer, r *http.Request) {
	fmt.Println(r.URL.String())
	fmt.Println(w.Write([]byte("HTTP/1.1 200 OK\nServer: nginx/1.26.0\nDate: Thu, 04 Jul 2024 02:12:54 GMT\nContent-Type: text/plain\nContent-Length: 7\nLast-Modified: Wed, 03 Jul 2024 08:08:59 GMT\nConnection: keep-alive\nETag: \"6685071b-7\"\nAccept-Ranges: bytes\n\n123456\n")))
}

// NewMiddlemanHttp 创建一个http中间人
func NewMiddlemanHttp(CaPath, CaKeyPath string) (*MiddlemanHttp, error) {
	middlemanHttp := &MiddlemanHttp{}

	ca, err := os.ReadFile(CaPath)
	if err != nil {
		return nil, err
	}

	key, err := os.ReadFile(CaKeyPath)
	if err != nil {
		return nil, err
	}

	loadCA, err := certificate.LoadCA(ca, key)
	if err != nil {
		return nil, err
	}
	middlemanHttp.CACert = loadCA
	middlemanHttp.HttpServer = http.Server{
		Handler: middlemanHttp,
		TLSConfig: &tls.Config{
			GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
				fmt.Println("sd")
				return BuildCa(middlemanHttp.CACert, info.ServerName)
			},
		},
	}

	return middlemanHttp, nil
}
