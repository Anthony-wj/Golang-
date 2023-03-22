/*
根据变量指向的值，是否是内存地址，我们将变量分为两种：

	普通变量：存数据值本身
	指针变量：存值的内存地址

& ：从一个普通变量中取得内存地址
* ：当*在赋值操作符（=）的右边，是从一个指针变量中取得变量值，当*在赋值操作符（=）的左边，是指该指针指向的变量
指针的零值是nil
*/
package main

import "fmt"

type demoPointer struct{}

func (demo demoPointer) demo01() {
	// 指针的创建：方法1
	// 先定义一个普通变量，然后通过&获取普通变量的地址，创建指针
	a := 1
	ptr := &a
	fmt.Printf("ptr: %v\n", ptr)
	fmt.Printf("p*tr: %v\n", *ptr)
}

func (demo demoPointer) demo02() {
	// 指针的创建：方法2
	// 先创建好指针，分配好内存后，再想变量所指向的地址写入相应的值
	ptr := new(int)
	*ptr = 1
	fmt.Printf("ptr: %v\n", ptr)
	fmt.Printf("*ptr: %v\n", *ptr)
}

func (demo demoPointer) demo03() {
	// 指针的创建：方法3
	// 先声明一个指针变量，然后向该变量指向的内存地址写入值
	var pint *int
	*pint = 1
	fmt.Printf("pint: %v\n", pint)
	fmt.Printf("*pint: %v\n", *pint)

}

func modifySlice(sls []int) {
	sls[0] = 10
}
func modifyPtr(ptr *[3]int) {
	(*ptr)[2] = 20
}

func main() {
	demo := demoPointer{}
	demo.demo01()
}
