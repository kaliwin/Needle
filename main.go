package main

import (
	"fmt"
	http2 "github.com/kaliwin/Needle/network/http"
)

type Tc struct {
	ClientId string `json:"clientId"`
	Data     string `json:"data"`
	Status   int    `json:"status"`
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
