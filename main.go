package main

import (
	"fmt"
	"github.com/kaliwin/Needle/ObjectConversion/ByteClassification/Classification"
	"golang.org/x/net/html"
	"io"
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

func main() {

	classification, err := Classification.NewByteClassification("/root/tmp/cci", "/root/tmp/body")
	if err != nil {
		panic(err)
	}

	err = classification.Do()
	fmt.Println(err)

}
