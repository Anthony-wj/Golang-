/*
	数组：由固定长度的特定类型元素组成的序列
	切片：可以容纳若干类型相同元素的容器，底层结构是数组
	切片是对数组的一个连续片段的引用，所以切片是一个引用类型
	切片构造的四种方法：
		1.从数组中截取arr[0:2]
		2.直接声明
		3.使用make函数构造 make(Type, size, cap)
		4.使用和数组一样偷懒的方法
	由于切片是引用类型，如果不对它复制的话，它的零值是nil
*/

package main

import "fmt"

type demoArray struct {
}

func (demo demoArray) demo01() {
	// 定义数组，但没有初始化
	var a [3]int
	fmt.Printf("a: %v\n", a)
}

func (demo demoArray) demo02() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Printf("a: %v\n", a)
}

func (demo demoArray) demo03() {
	// 自动推断数组长度
	var a = [...]int{1, 2, 3}
	fmt.Printf("a: %v\n", a)
}

func (demo demoArray) demo04() {
	type array5 [5]int
	var a array5 = array5{1, 2, 3, 4, 5}
	fmt.Printf("a: %v\n", a)
}

func (demo demoArray) demo05() {
	var a [5]int = [5]int{2: 10}
	fmt.Printf("a: %v\n", a)
}

type demoSlice struct {
}

func (demo demoSlice) demo01() {
	// 切片的获取方法1
	a := [...]int{1, 2, 3, 4, 5}
	s := a[1:4] // 切片左闭右开
	fmt.Printf("s: %T %v\n", s, s)
}

func (demo demoSlice) demo02() {
	// 定义切片，但不初始化
	var a []string
	fmt.Printf("a: %v\n", a)
	// 定义并初始化切片
	var b []int = []int{1, 2, 3}
	fmt.Printf("b: %v\n", b)
	// 使用make函数
	c := make([]int, 2)
	fmt.Println(c, len(c), cap(c))
	d := make([]int, 2, 10)
	fmt.Println(d, len(d), cap(d))
	// 指定下标赋值
	e := []int{4: 2}
	fmt.Printf("c: %v\n", e)
}

func (demo demoSlice) demo03() {
	myarr := []int{1}
	// 增加一个元素
	myarr = append(myarr, 2)
	// 增加多个元素
	myarr = append(myarr, 3, 4)
	// 增加一个切片 ...表示解包
	myarr = append(myarr, []int{7, 8}...)
	myarr = append([]int{0}, myarr...)
	myarr = append(myarr[:5], append([]int{5}, myarr[5:]...)...)
	myarr = append(myarr[:6], append([]int{6}, myarr[6:]...)...)
	fmt.Printf("myarr: %v\n", myarr)
}

func (demo demoSlice) demo04() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s := arr[4:6:8]
	fmt.Printf("s: %v\n", s)
	fmt.Printf("cap(s): %v\n", cap(s))
	s = s[:cap(s)]
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("s[1]: %v\n", cap(s))
}

func main() {
	demo := demoSlice{}
	demo.demo04()

	// myarrr := []int{1}
	// // 追加一个元素
	// myarrr = append(myarrr, 2)
	// // 追加多个元素
	// myarrr = append(myarrr, 3, 4)
	// // 追加一个切片, ... 表示解包，不能省略
	// myarrr = append(myarrr, []int{7, 8}...)
	// // 在第一个位置插入元素
	// myarrr = append([]int{0}, myarrr...)
	// // 在中间插入一个切片(两个元素)
	// myarrr = append(myarrr[:5], append([]int{5, 6}, myarrr[5:]...)...)
	// fmt.Println(myarr)

}
