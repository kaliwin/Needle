package main

import (
	"bufio"
	"fmt"
	http2 "github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/http"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

func isLegalHTML(htmlContent string) (bool, error) {
	tokenizer := html.NewTokenizer(strings.NewReader(htmlContent))

	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			// 检查是否有错误
			if tokenizer.Err() != io.EOF {
				return false, fmt.Errorf("HTML is not well-formed: %v", tokenizer.Err())
			}
			break
		}
	}
	return true, nil
}

type Hu struct {
	re *bufio.Reader
}

func (h Hu) Read(p []byte) (n int, err error) {
	return h.re.Read(p)
}

func (h Hu) Close() error {
	return nil
}

func main() {

	request, _ := http.NewRequest("GET", "http://www.google.com/asdf?sdf=sdf", nil)

	//all := &http2.ReadAll{}
	//
	//err := request.Write(all)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf(string(all.Data))

	httpRequest, err := http2.ParseHttpRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(httpRequest.GetPath())

}
