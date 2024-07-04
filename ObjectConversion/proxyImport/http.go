package proxyImport

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"github.com/kaliwin/Needle/IO/Interface"
	"github.com/kaliwin/Needle/IO/file"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	http2 "github.com/kaliwin/Needle/PublicStandard/mark/http"
	"github.com/kaliwin/Needle/crypto/certificate"
	http3 "github.com/kaliwin/Needle/network/http"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
)

// 纯http代理方式导入

// HttpProxyImport http 代理导入
type HttpProxyImport struct {
	// 代理地址
	//TarGetProxy http3.HttpClient // 目标代理地址
	// 原始数据目录
	httpGroupPath string // 原始数据目录 目录文件必须为httpGroup 只支持integrate的onlyHost格式目录 一个host的所有请求放在一个文件夹下
	// 不允许使用path分成 那样会导致索引丢失

	CACert certificate.CACert // CA证书 用于签发证书 通常burp可以设置信任所有证书
	//Client http3.HttpClient   // http客户端
}

// StartHttpServer 启动http服务器
func (h *HttpProxyImport) StartHttpServer(addr string) error {
	return http.ListenAndServe(addr, h)
}

// ServeHTTP http代理服务器 处理代理请求
func (h *HttpProxyImport) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	// 根据请求方法分发处理函数
	if request.Method == http.MethodConnect {
		h.handleHTTPS(writer, request)
	} else {
		h.handleHTTP(writer, request)
	}
}

// handleHTTP http的处理
func (h *HttpProxyImport) handleHTTP(w http.ResponseWriter, r *http.Request) {
	hijacker, _ := w.(http.Hijacker)
	hijack, _, err := hijacker.Hijack() // 获取底层连接

	if err != nil {
		return
	}
	defer func() { _ = hijack.Close() }()
	h.handleResponse(hijack, r)
}

// handleHTTPS https的处理
func (h *HttpProxyImport) handleHTTPS(w http.ResponseWriter, r *http.Request) {
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
	server := tls.Server(hijack, &tls.Config{ // 设置tls配置 强制使用http1.1
		GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
			return BuildCa(h.CACert, info.ServerName) // 使用CA签发证书
		},
		NextProtos: []string{"http/1.1"}, // 强制使用http1.1 TLS协议保持默认
	})
	defer func() { _ = server.Close() }() // 关闭连接 因为是http1.1 所以每次请求都会关闭连接

	request, err := http.ReadRequest(bufio.NewReader(server)) // 读取请求
	if err != nil {
		log.Println(err)
		return
	}
	h.handleResponse(server, request)
}

// handleResponse 处理响应
func (h *HttpProxyImport) handleResponse(w io.Writer, r *http.Request) {
	httpGroupId := r.Header.Get(string(http2.MarkHttpGroupID)) // 获取http组ID

	if httpGroupId == "test" {
		_, err := w.Write([]byte("HTTP/1.1 200\nDate: Wed, 03 Jul 2024 10:26:33 GMT\nLength: 1223\nContent-Length: 0\n\n"))
		if err != nil {
			log.Panicln(err)
			return
		}
		return
	}

	if len(httpGroupId) > 1 {
		fileC := fmt.Sprintf("%s.httpGroup", httpGroupId)
		readFile, err := os.ReadFile(h.httpGroupPath + "/" + fileC) // 读取文件
		if err != nil {
			log.Panicln(err)
			return
		}
		res := HttpStructureStandard.HttpReqAndRes{}
		err = proto.Unmarshal(readFile, &res)
		if err != nil {
			log.Panicln(err)
			return
		}
		_, err = w.Write(res.GetRes().GetData()) // 写入数据
		if err != nil {
			log.Panicln(err)
			return
		}
	}
}

// Go 启动http代理导入
func (h *HttpProxyImport) Go(tarGetProxy string) {

	parse, _ := url.Parse(tarGetProxy)

	client := http.Client{
		Transport: &http.Transport{
			Proxy: func(r *http.Request) (*url.URL, error) {
				return parse, nil
			},
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	stream, err := file.BuildFIleObjectStream(h.httpGroupPath, true, Interface.ObjectTypeHttpGroupProto)
	if err != nil {
		log.Println(err)
		return
	}

	taskPool := make(chan *HttpStructureStandard.HttpReqAndRes, 100)

	group := sync.WaitGroup{}

	for i := 0; i < 20; i++ { // 10个工人

		go func() {
			group.Add(1)
			defer group.Done()
			for res := range taskPool { // 取任务
				request, err := http3.BuildRequest(res.GetReq()) // 转为标准请求
				if request == nil {
					log.Println(err)
					continue
				}

				request.Header.Set(string(http2.MarkHttpGroupID), res.GetInfo().GetId()) // 设置http组ID
				request.RequestURI = ""

				request.URL.Host = fmt.Sprintf("%s:%d", res.GetReq().GetHttpReqService().GetIp(), res.GetReq().GetHttpReqService().GetPort())

				if res.GetReq().GetHttpReqService().GetSecure() {
					request.URL.Scheme = "https"
				} else {
					request.URL.Scheme = "http"
				}
				_, err2 := client.Do(request) // 通过代理发送请求
				if err2 != nil {
					log.Println(err2)
					continue
				}

			}
		}()

	}

	stream.Iteration(func(a any) bool {
		if res, ok := a.(*HttpStructureStandard.HttpReqAndRes); ok {

			taskPool <- res

		}
		return true
	})

	close(taskPool)
	group.Wait()

}

// BuildHttpProxyImport 构建http代理导入
func BuildHttpProxyImport(httpGroupPath, caPath, caKeyPath string) (HttpProxyImport, error) {
	proxyImport := HttpProxyImport{}

	ca, err := os.ReadFile(caPath)
	if err != nil {
		return HttpProxyImport{}, err
	}
	caKey, err := os.ReadFile(caKeyPath)
	if err != nil {
		return HttpProxyImport{}, err
	}

	pemCA, err := certificate.LoadCA(ca, caKey)
	if err != nil {
		return HttpProxyImport{}, err
	}

	proxyImport.httpGroupPath = httpGroupPath
	proxyImport.CACert = pemCA

	return proxyImport, nil
}

// BuildCa 签发证书
func BuildCa(CA certificate.CACert, name string) (*tls.Certificate, error) {
	cert, privacy, err := certificate.MakePemCert(CA, []string{"*." + name, name}, name, certificate.GetTlsCASubject())
	if err != nil {
		return nil, err
	}
	pair, err := tls.X509KeyPair(cert, privacy)

	return &pair, err
}

//////
/////
//////
/////
//////
///////////
///////
/////
//////
/////
//////
///////////
////////////
////////
///////
////////
/////////////
//////////////
////////
///////
////////
/////////////
//////////////
////////
///////
////////
/////////////
//////////////
////////
///////
////////
/////////////
//////////////
////////
///////
////////
/////////////
//////////////
////////
///////
////////
/////////////
/////////
