package main

import (
	"fmt"
	"github.com/dop251/goja"
	"log"
)

func main() {
	// 一个简单的 JavaScript 代码示例
	javascriptCode := `
	function add(ax, b) {
	var x = 10;
	subtract(ax,"c");
	   return a + b;
	}
	
	function subtract(ac, b) {
		// 测试
	   return a - b;
	}
	
	var x = 10;
	var y = 20;
	var c  = add(x, y);
	
	`

	//javascriptCode, _ := os.ReadFile("/root/tmp/app.js")

	prg, err := goja.Parse("app.js", string(javascriptCode))
	if err != nil {
		log.Println(err)
	}

	fmt.Println(prg)

	//JsStruct := js.JsStruct{}
	//
	//JsStruct.ParseJsAST(prg)
	//
	//fmt.Println(JsStruct)

}
