package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestParseHttpRequest(t *testing.T) {

	request, _ := http.NewRequest("GET", "https://www.baidu.com", nil)

	httpRequest, err := ParseHttpRequest(request)
	if err != nil {
		t.Error(err)
	}

	//httpRequest.AddPath("ji")
	httpRequest.AddQueryValue("name", "jiji")
	httpRequest.AddQueryValue("age", "23 aasdf and 1=1 s")

	httpRequest.SetHeadValue("Host", "sd")

	httpRequest.SetBody([]byte("hello world"))

	httpRequest.DeleteQuery("age")
	httpRequest.SetMethod("PUT")
	//httpRequest.AddHead("Accept", "*/*")
	httpRequest.DeleteHead("Accept")
	message := httpRequest.BuildMessage()
	fmt.Println(httpRequest.GetUrl().String())
	fmt.Println(string(message))

}
