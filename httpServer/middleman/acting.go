package middleman

import (
	"bufio"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"github.com/kaliwin/Needle/IO/Interface"
	"github.com/kaliwin/Needle/IO/file"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"github.com/kaliwin/Needle/crypto/certificate"
	http2 "github.com/kaliwin/Needle/network/http"
	"log"
	"net/http"
	"net/url"
	"os"
)

// handleHttps https 的隧道处理 强制使用http1.1 因为http2是长连接 而现在拿到的是net.conn低级连接
// 处理http2协议 过于复杂 未来将考虑从其他开源库中寻求http2的解决方案
func (receiver Acting) handleHttps(w http.ResponseWriter, r *http.Request) {
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	hijack, _, err := hijacker.Hijack() // 获取底层连接
	if err != nil {
		return
	}
	_, err = hijack.Write([]byte("HTTP/1.0 200 Connection established\r\n\r\n")) // 返回200
	if err != nil {
		return
	}

	server := tls.Server(hijack, &tls.Config{ // 设置tls配置 强制使用http1.1
		GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
			return BuildCa(receiver.CACert, info.ServerName) // 使用CA签发证书
		},
		NextProtos: []string{"http/1.1"}, // 强制使用http1.1 TLS协议保持默认
	})
	defer func() { _ = server.Close() }() // 关闭连接 因为是http1.1 所以每次请求都会关闭连接

	request, err := http.ReadRequest(bufio.NewReader(server)) // 读取请求
	if err != nil {
		fmt.Println("[-] sdsdsdsd")
		log.Println(err)
		return
	}
	fmt.Println(request.URL.String())
	request.RequestURI = ""
	request.URL.Scheme = "https"
	request.URL.Host = request.Host
	do, err := receiver.Client.Do(request) // 发送请求
	if err != nil {
		//log.Println(err)
		return
	}

	err = do.Write(server) // 写回响应
	if err != nil {
		return
	}

}

// handleHTTP http 的处理
func (receiver Acting) handleHTTP(w http.ResponseWriter, r *http.Request) {
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	hijack, _, err := hijacker.Hijack() // 获取底层连接
	if err != nil {
		return
	}

	r.RequestURI = ""
	resp, err := receiver.Client.Do(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	_ = resp.Write(hijack) // 将原始报文直接写回
}

func (receiver Acting) handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	// 根据请求方法分发处理函数
	if r.Method == http.MethodConnect {

		receiver.handleHttps(w, r)
	} else {

		receiver.handleHTTP(w, r)
	}
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

func (receiver Acting) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	receiver.handleRequestAndRedirect(writer, request)

}

// BuildActing 构建一个演员通过上游代理处理请求
func BuildActing(proxyPort string, caPath string, keyPath string) (*Acting, error) {

	burpCA, err := os.ReadFile(caPath)
	if err != nil {
		return nil, nil
	}
	burpCAKey, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, nil
	}

	ca, er := certificate.LoadCA(burpCA, burpCAKey) // 加载CA
	if er != nil {
		return nil, nil
	}
	var client http.Client

	if proxyPort == "" {
		client = http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	} else {

		parse, err := url.Parse(proxyPort)
		if err != nil {
			return nil, err
		}

		client = http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				Proxy: func(r *http.Request) (*url.URL, error) {
					return parse, nil
				},
			},
		}
	}
	return &Acting{
		CACert: ca,
		Client: &client,
	}, nil
}

// Acting 演员 模拟一个正常代理服务器 用于监听和篡改流量
type Acting struct {
	CACert certificate.CACert // CA证书
	Client http2.HttpClient
}

// BuildObjectC 构建对象转换
func BuildObjectC() {

	burpCA, err := os.ReadFile("/root/cyvk/ManDown/CA/burpCA.cer")
	if err != nil {
		return
	}
	burpCAKey, err := os.ReadFile("/root/cyvk/ManDown/CA/burpCA-key.cer")
	if err != nil {
		return
	}

	ca, _ := certificate.LoadCA(burpCA, burpCAKey) // 加载CA

	acting := Acting{
		CACert: ca,
		Client: Hu{},
	}
	server := http.Server{
		Addr:    ":8081",
		Handler: acting,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

type Hu struct {
}

func (h Hu) Do(req *http.Request) (*http.Response, error) {
	hc := HW{}
	err := req.Write(&hc)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(hc.data))
	return nil, nil
}

type HW struct {
	data []byte
}

func (H *HW) Write(p []byte) (n int, err error) {
	H.data = append(H.data, p...)
	return len(H.data), nil
}

type Yi struct {
	flow ResFlow
	Lock chan bool
}

type ResFlow struct {
	ReqSign string
	ResData *HttpStructureStandard.HttpResData
}

func (y *Yi) Go() {
	hp, _ := url.Parse("http://localhost:8080")
	burpProxy := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy: func(r *http.Request) (*url.URL, error) {
				return hp, nil
			},
		},
	}
	stream, err := file.BuildFIleObjectStream("/root/tmp/data", true, Interface.ObjectTypeHttpGroupListProto)

	if err != nil {
		panic(err)
	}

	stream.Iteration(func(a any) bool {
		list := a.(*HttpStructureStandard.HttpRawByteStreamList)
		for _, res := range list.GetHttpRawByteStreamList() {

			if len(res.Req.GetData()) < 1 || len(res.Res.GetData()) < 1 {
				continue
			}

			sum := md5.Sum(res.GetReq().GetData())
			he := hex.EncodeToString(sum[:])

			req, err2 := http2.BuildRefactorStandardHttpReq(res.GetReq(), &burpProxy)
			if err2 != nil {
				log.Println(err2)
				continue
			}

			//y.Lock <- true
			y.flow = ResFlow{ // 写入流
				ReqSign: he,
				ResData: res.GetRes(),
			}
			_, err2 = req.Send() // 发送请求
			if err2 != nil {
				log.Println(err2)
				continue
			}

		}
		fmt.Println("Done")
		return true
	})

}

func (y *Yi) Do(req *http.Request) (*http.Response, error) {

	hc := HW{}
	err := req.Write(&hc)
	if err != nil {
		panic(err)
	}
	sum := md5.Sum(hc.data)
	he := hex.EncodeToString(sum[:])

	if y.flow.ReqSign == he {
		res, err := http2.BuildRefactorStandardHttpRes(y.flow.ResData, nil)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return res.BuildResponse()
	}
	fmt.Println("sdfasfd")
	return nil, nil
}
