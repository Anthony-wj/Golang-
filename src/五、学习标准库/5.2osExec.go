/*
在Golang中用于执行命令的库是os/exec,exec.Command函数返回一个Cmd对象，根据不同的需求，可以将命令的执行分为三种情况：
	1.只执行命令，不获取结果
	2.执行命令，并获取结果(不区分stdout和stderr)
	3.执行命令，并获取结果(区分stdout和stderr)
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type demoOs struct{}

func (demo demoOs) demo01_Run() {
	cmd := exec.Command("ls", "-l", "/var/log/")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func (demo demoOs) demo02_CombinedOutput() {
	cmd := exec.Command("ls", "-l", ".")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n %s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s \n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

func (demo demoOs) demo03_Std() {
	cmd := exec.Command("ls", "-l", "~")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

}

func main() {
	demo := demoOs{}
	demo.demo03_Std()
}
