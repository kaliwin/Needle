package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义命令行参数
	var name string
	var age int
	var verbose bool
	//var verName = ""

	flag.StringVar(&name, "name", "Guest", "Specify your name")
	flag.IntVar(&age, "age", 0, "Specify your age")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose mode")
	//flag.StringVar(&verName, "verbose", "", " test ")

	// 解析命令行参数
	flag.Parse()

	// 使用解析后的参数
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Verbose:", verbose)
	//fmt.Println("verName:", verName)
}
