package main

import (
	"fmt"
	"flag"
	"os"
)

/**
命令行输入参数

demo
go run demo1.go -name 正东
 */
var name1 string

var cmdLine = flag.NewFlagSet("question", flag.ExitOnError) // flag.PanicOnError 出错的时候打什么样子错误

func init() {
	cmdLine.StringVar(&name1, "name", "everyone", "The greeting object")
	//flag.StringVar(&name1, "name", "everyone", "The greeting object")
}

func main() {
	/*	flag.Usage = func() {
			fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
			flag.PrintDefaults()
		}*/
	/*	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
		flag.CommandLine.Usage = func() {
			fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
			flag.PrintDefaults()
		}*/

	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello,%s!\n", name1)
}
