package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kaliwin/Needle/MorePossibilityApi/grpc/BurpMorePossibilityApi"
	http2 "github.com/kaliwin/Needle/network/http"
	"io"
	"log"
	"net/http"
)

type Tc struct {
	ClientId string `json:"clientId"`
	Data     string `json:"data"`
	Status   int    `json:"status"`
}

type test struct {
}

// IntruderPayloadProcessor burp迭代载荷处理器
func (t test) IntruderPayloadProcessor(ctx context.Context, data *BurpMorePossibilityApi.PayloadProcessorData) (*BurpMorePossibilityApi.ByteData, error) {

	phone := string(data.GetPayload())

	resp, err := http.Get("http://127.0.0.1:5612/business-demo/invoke?group=test&action=encrypt&phone=" + phone)
	if err != nil {
		log.Println(err)
		return &BurpMorePossibilityApi.ByteData{ByteData: data.GetPayload()}, nil
	}
	d, _ := io.ReadAll(resp.Body)
	tc := Tc{}
	json.Unmarshal(d, &tc)
	return &BurpMorePossibilityApi.ByteData{ByteData: []byte(tc.Data)}, nil
}

func main() {
	//// 定义命令行参数
	//var name string
	//var age int
	//var verbose bool
	////var verName = ""
	//
	//flag.StringVar(&name, "name", "Guest", "Specify your name")
	//flag.IntVar(&age, "age", 0, "Specify your age")
	//flag.BoolVar(&verbose, "verbose", false, "Enable verbose mode")
	////flag.StringVar(&verName, "verbose", "", " test ")
	//
	//// 解析命令行参数
	//flag.Parse()
	//
	//// 使用解析后的参数
	//fmt.Println("Name:", name)
	//fmt.Println("Age:", age)
	//fmt.Println("Verbose:", verbose)
	////fmt.Println("verName:", verName)

	//burpServer, err := MorePossibilityApi.NewBurpGrpcServer(":9000")
	//if err != nil {
	//	log.Println(err)
	//}
	//burpServer.RegisterIntruderPayloadProcessorServer(&test{})
	//
	//err = burpServer.Start()
	//if err != nil {
	//	log.Println(err)
	//}

	fmt.Println(http2.DefaultHeader["sd"])
	//
	//client := BurpMorePossibilityApi.NewBurpServerClient(nil)
	//client.

}
