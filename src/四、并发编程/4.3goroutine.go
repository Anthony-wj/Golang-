/*
	一个goroutine本身就是一个函数，当你直接调用时，他就是一个普通函数，如果你在调用前加一个关键字go，那你就开启了一个goroutine。
	1.协程的初步使用
		一个Go程序的入口通常是main函数，程序启动后，main函数最先运行，我们称之为main goroutine
		在main中或者其下嗲用的代码中才可以使用go + func()的方法来启动协程
		main的地位相当于主线程，当main函数执行完成后，这个县城也就终结了，其下的运行着的所有协程也不管代码是不是还在跑，也得乖乖退出
		time.Sleep(time.Second)阻塞main协程
*/

package main

import (
	"fmt"
	"time"
)

type demoGoroutine struct{}

func (demo demoGoroutine) demo01() {
	fmt.Println("hello, go")
}

func (demo demoGoroutine) demo02(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("name: %v\n", name)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	demo := demoGoroutine{}
	go demo.demo02("协程1号")
	go demo.demo02("协程2号")
	time.Sleep(time.Second)
}
