package main

import (
	"flag"
	"fmt"
	"os"
)

type demoFlag struct{}

func (demo demoFlag) demo01_Args() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}

func (demo demoFlag) demo02_StringVar() {
	var name string
	// flag.StringVar定义了一个字符串参数，它接收几个参数
	// func StringVar(p *string, name string, value string, usage string)
	flag.StringVar(&name, "name", "jack", "your name")
	flag.Parse()
	fmt.Printf("name: %v\n", name)
}

var name string

func init() {
	flag.StringVar(&name, "name", "jack", "your name")
}

func (demo demoFlag) demo02_Plus() {
	flag.Parse()
	fmt.Printf("name: %v\n", name)
}

func main() {
	demo := demoFlag{}
	demo.demo02_Plus()
}
