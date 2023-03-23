/*
	1.关于函数
		函数是基于功能或逻辑的可复用的代码结构。将一段功能复杂、很长的一段代码封装
		成多个代码片段(即函数)，有助于提高代码可读性和可维护性。

		在Go语言中，函数可以分为两种：
			带有名字的普通函数
			没有名字的匿名函数

		由于Go语言是编译型语言，所以函数编写的顺序是无关紧要的，它不像python那样，
		函数在位置上需要定义在调用之前。

	2.函数的声明
		函数的声明，使用关键字func，后面一次接函数名、参数列表、返回值列表，用{}包裹的代码逻辑体
			func 函数名(形式参数列表)(返回值列表){
				函数体
			}
		形式参数列表描述了函数的参数名以及参数类型，这些参数作为局部变量，其值由函数调用者提供。
		返回值列表描述了函数返回值的变量名以及类型，如果函数返回一个无名变量或者没有返回值，
		返回值列表的括号是可以省略的。

	3.函数实现可变参数
		可变参数分为两类：
			多个类型一直的参数  (args ...int)
			多个类型不一致的参数 (args ...interface{})

	4.多个可变参数函数传递参数
		上面提到了可以使用 ... 来接收多个参数，除此之外，它还有一个用法，就是用来解序列，
		将函数的可变参数（一个切片）一个一个取出来，传递给另一个可变参数的函数，而不是传递可变参数变量本身。

	5.函数的返回值
		可以没有返回值，函数体可以用return来结束运行
		可以返回多个值
		可以返回带有变量的值

	6.方法与函数
		方法是一种特殊的函数。当啊一个函数和对象/结构体进行绑定的时候，我们就称这个函数是一个方法

	7.匿名函数的使用
		懒得定义函数名
		回调函数使用

*/

package main

import "fmt"

type demoFunc struct{}

func (demo demoFunc) demo01(a, b int) (c int) {
	c = a + b
	return
}

func (demo demoFunc) demo02(args ...int) (sum int) {
	for _, v := range args {
		sum += v
	}
	return
}

func (demo demoFunc) demo03(args ...interface{}) {
	for _, arg := range args {
		fmt.Printf("arg: %v\n", arg)
	}
}
func (demo demoFunc) demo04_sum(args ...int) (sum int) {
	for _, v := range args {
		sum += v
	}
	return
}
func (demo demoFunc) demo04_Sum(args ...int) int {
	result := demo.demo04_sum(args...)
	return result
}

func (demo demoFunc) demo05(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func main() {
	demo := demoFunc{}
	result := demo.demo04_Sum(1, 2)
	fmt.Printf("result: %v\n", result)
	demo.demo05([]int{1, 2, 3, 4, 5}, func(v int) {
		fmt.Println(v)
	})
}
