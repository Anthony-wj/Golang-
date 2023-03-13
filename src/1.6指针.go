/*
根据变量指向的值，是否是内存地址，我们讲变量分为两种：

	普通变量：存数据值本身
	指针变量：存值的内存地址

& ：从一个普通变量中取得内存地址
* ：当*在赋值操作符（=）的右边，是从一个指针变量中取得变量值，当*在赋值操作符（=）的左边，是指该指针指向的变量
指针的零值是nil
*/
package main

import "fmt"

func modifySlice(sls []int) {
	sls[0] = 10
}
func modifyPtr(ptr *[3]int) {
	(*ptr)[2] = 20
}

func main() {
	// 指针的创建：
	// 方法1：先定义对应的变量，再通过变量取得内存地址，创建指针。
	// aint := 1
	// prt := &aint

	// 方法2：先创建指针，分配好内存后，再给指针指向的内存地址写入对应的值。
	// astr := new(string)
	// *astr = "hello world"

	// 方法3：先声明一个指针变量，再从其他变量取得内存地址赋值给它。
	// var bint *int
	// bint = &aint

	a := [3]int{1, 2, 3}
	modifySlice(a[:])
	fmt.Printf("a: %v\n", a)
	modifyPtr(&a)
	fmt.Printf("a: %v\n", a)
}
