package main

import (
	"fmt"
	"github.com/kaliwin/Needle/MorePossibilityApi"
	"github.com/kaliwin/Needle/MorePossibilityApi/grpc/BurpMorePossibilityApi"
	"log"
	"os"
	"time"
)

func main() {

	server, err := MorePossibilityApi.NewGrpcServer(":1080")
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	BurpMorePossibilityApi.RegisterHttpFlowHandlerServer(server.Server, MorePossibilityApi.ApiTest{})
	server.Start()
	fmt.Println(server.GetStatus())
	time.Sleep(time.Hour * 1)
}
