package OriginalMessage

import (
	"fmt"
	"testing"
)

func TestRequestOriginalMessage(t *testing.T) {

	head := make(Head)

	head[Host] = "ss0.baidu.com"
	head[UA] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) "
	head[Accept] = "*/*"
	//body := []byte("{\"passWord\":\"TKVlWsnub5L9bPVXsPIVrvYtxjEVD8sP\",\"userName\":\"cyvk\",\"validateCode\":\"7275\"}")

	//head[ContentLength] = fmt.Sprintf("%d", len(body))
	message := RequestOriginalMessage{
		RequestLine: RequestLine{
			Method:      "GET",
			Path:        "/6ONWsjip0QIZ8tyhnq/ps_default.gif?_t=1720347026928",
			HttpVersion: Http1,
		},
		Head: head,
		Body: nil,
	}

	buildMessage := message.BuildMessage()
	fmt.Println(string(buildMessage))
	originalMessage, err := ParseRequestOriginalMessage(buildMessage)
	if err != nil {
		panic(err)
	}
	fmt.Println(originalMessage)
}
