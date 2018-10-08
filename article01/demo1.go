package main

import (
	"fmt"
	"flag"
)

/**
命令行输入参数

demo
go run demo1.go -name 正东
 */
var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello,%s!\n", name)
}

