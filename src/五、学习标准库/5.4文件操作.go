/*
整个文件读入内存
	1.直接指定文件名读取
		os.ReadFile
		ioutil.ReadFile
		两者作用一样

	2.先创建句柄再读取
		os.Open

每次读取一行
	1.bufio.ReadLine()
	2.bufio.ReadBytes()
	3.bufio.ReadString()

每次读取固定字节数
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type demoRead struct{}

func (demo demoRead) demo01() {
	content, err := os.ReadFile("d:/code/golang/Golang-Codes/src/五、学习标准库/diff.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(content))
	for n, line := range content {
		fmt.Printf("%v line: %s\n", n, string(line))

	}
}

func (demo demoRead) demo02() {
	fi, err := os.Open("d:/code/golang/Golang-Codes/src/五、学习标准库/diff.txt")
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(fi)
	for {
		lineBytes, err := r.ReadBytes('\n')
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
}

func main() {
	demo := demoRead{}
	demo.demo02()
}
