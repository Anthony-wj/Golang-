/*
	1.什么是空接口？
		空接口是特殊形式的接口类型，普通的接口都有方法，而空接口没有定义任何方法，也因此，我们可以说所有类型都至少实现了空接口。

		每个接口都包含两个属性，一个是值，一个是类型。
		而对于空接口来说，这两者都是nil

	2.如何使用空接口？
		第一，通常我们会直接使用 interface{} 作为类型声明一个实例，而这个实例可以承载任意类型的值。

		第二，如果想让你的函数可以接收任意类型的值 ，也可以使用空接口
			func myfunc(iface interface{})
			func myfunc(ifaces ...interface{})

	3.空接口几个要注意的坑
		坑1：空接口可以承载任意值，但不代表任意类型就可以承接空接口类型的值
		坑2：当空接口承载数组和切片后，该对象无法再进行切片
		坑3：当你使用空接口来接收任意类型的参数时，它的静态类型是 interface{}，但动态类型（是 int，string 还是其他类型）我们并不知道，因此需要使用类型断言。

*/

package main

import (
	"fmt"
)

func myfunc(iface interface{}) {
	fmt.Printf("iface: %v\n", iface)
}

func myfunc2(ifaces ...interface{}) {
	for i, v := range ifaces {
		fmt.Println(i, v)
	}
}

func main() {
	i := 1
	s := "s"
	b := true
	myfunc(i)
	myfunc(s)
	myfunc(b)
	myfunc2(i, s, b)

	any := make([]interface{}, 5)
	// value, ok := any.(int)
	// fmt.Println(value, ok)
	any[0] = 18
	any[1] = "hello world"
	any[2] = []int{1, 2, 3, 4}
	for _, value := range any {
		fmt.Printf("value: %v\n", value)
	}
}
