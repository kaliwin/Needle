package main

import (
	"context"
	"github.com/kaliwin/Needle/MorePossibilityApi/grpc/BurpMorePossibilityApi"
	"github.com/kaliwin/Needle/httpServer/middleman"
	"log"
)

type Test struct {
}

func (t Test) HttpHandleRequestReceived(ctx context.Context, data *BurpMorePossibilityApi.HttpFlowReqData) (*BurpMorePossibilityApi.HttpRequestAction, error) {
	if data.GetHttpFlowSource() == BurpMorePossibilityApi.HttpFlowSource_REPEATER {
		//client := http.Client{}
		//
		//proxy, _ := url.Parse("http://127.0.0.1:8080")
		//
		//client.Transport = &http.Transport{
		//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		//	Proxy: http.ProxyURL(proxy),
		//}
		//
		//req, err := StandardHttp.BuildRefactorStandardHttpReq(data.GetHttpReqGroup().GetHttpReqData(), &client)
		//if err != nil {
		//	log.Println(err)
		//	return &BurpMorePossibilityApi.HttpRequestAction{Continue: true}, nil
		//}
		//
		//

	}
	return &BurpMorePossibilityApi.HttpRequestAction{Continue: true}, nil
}

func (t Test) HttpHandleResponseReceived(ctx context.Context, data *BurpMorePossibilityApi.HttpFlowResData) (*BurpMorePossibilityApi.HttpResponseAction, error) {

	return &BurpMorePossibilityApi.HttpResponseAction{Continue: true}, nil
}

func main() {

	err := middleman.StartMiddleman(":443", "http://127.0.0.1:8080", "/root/tmp/burpCA.cer", "/root/tmp/burpCA-key.cer")
	if err != nil {
		log.Println(err)
	}

	//MorePossibilityApi.ApiTest{}

	//
	//fmt.Println(fmt.Sprintf("%s:%d", "14.13.54.11", 56))
	//
	////client := http.Client{}
	//////proxy, _ := url.Parse("http://127.0.0.1:8080")
	////
	////client.Transport = &http.Transport{
	////	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	////	//Proxy:           http.ProxyURL(proxy),
	////}
	////
	////request, _ := http.NewRequest("GET", "https://www.baidu.com/dsi", nil)
	////
	////parse, _ := url.Parse("/")
	////parse.Scheme = "http"
	////parse.Host = "erin.server:9988"
	////
	////request.URL = parse
	////request.Host = "erin.server:sd"
	////do, err := client.Do(request)
	////if err != nil {
	////	log.Println(err)
	////	return
	////}
	////fmt.Println(do.Status)
	////
	////all, _ := io.ReadAll(do.Body)
	////fmt.Println(string(all))
	//
	////burpServer, err := MorePossibilityApi.NewGrpcServer(":9000")
	////if err != nil {
	////	log.Println(err)
	////	os.Exit(0)
	////}
	////
	////burpServer.RegisterHttpFlowHandlerServer(Test{})
	////
	////burpServer.Start()
	////fmt.Println(burpServer.GetStatus())
	////time.Sleep(time.Hour * 1)
	//
	//c, _ := os.ReadFile("/root/tmp/c.cer")
	//privDER, _ := os.ReadFile("/root/tmp/k.cer")
	//
	//memory := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: c})
	//toMemory := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privDER})
	//
	//os.WriteFile("/root/tmp/c.pem", memory, 0644)
	//os.WriteFile("/root/tmp/k.pem", toMemory, 0644)

}
