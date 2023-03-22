/*
	1. 延迟调用
	2. 即时求值的变量快照
		使用 defer 只是延时调用函数，此时传递给函数里的变量，不应该受到后续程序的影响。
		如果 defer 后面跟的是匿名函数，情况会有所不同， defer 会取到最后的变量值
	3.多个defer反序调用
		可见 多个defer 是反序调用的，有点类似栈一样，后进先出。
	4.defer 与 return 孰先孰后
		defer 在 return后调用
*/

package main

import "fmt"

type demoDefer struct{}

func (demo demoDefer) demo01() {
	defer fmt.Println("end")
	fmt.Println("start")
}

func myfunc() {
	fmt.Println("B")
}

func main() {
	demo := demoDefer{}
	demo.demo01()
	// defer myfunc()
	// fmt.Println("A")

	name := "go"
	// defer fmt.Println(name)
	defer func() {
		fmt.Println(name)
	}()
	name = "python"
	fmt.Println(name)
}
