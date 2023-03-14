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

func main() {
	// 数组定义 var a [3]int
	var arr0 [3]int
	fmt.Printf("arr0: %v\n", arr0)
	// 数组定义并初始化
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Printf("arr1: %v\n", arr1)
	arr2 := [3]int{1, 2, 3}
	fmt.Printf("arr2: %v\n", arr2)
	// [...]自动推断数组长度，arr3和arr4是不同类型的数组
	arr3 := [...]int{1, 2, 3}
	arr4 := [...]int{1, 2, 3, 4}
	fmt.Printf("arr3: %T\n", arr3)
	fmt.Printf("arr4: %T\n", arr4)

	// 觉得[3]int写起来麻烦，可以使用type关键字来定义一个类型字面量
	type arr5 [3]int
	myarr := arr5{1, 2, 3}
	fmt.Printf("%d 的类型是： %T", myarr, myarr)

	// 数组定义偷懒方法
	arr6 := [4]int{2: 3}
	fmt.Printf("arr6: %v\n", arr6)

	myarr2 := [...]int{1, 2, 3}
	fmt.Printf("%d 的类型是%T\n", myarr2, myarr2)
	fmt.Printf("%d 的类型是%T\n", myarr2[0:2], myarr2[0:2])

	var slice1 []string
	fmt.Printf("slice1: %v\n", slice1)

	var slice2 = []int{}
	fmt.Printf("slice2: %v\n", slice2)

	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))

	c := []int{4: 2}
	fmt.Printf("c: %v\n", c)

	myarrr := []int{1}
	// 追加一个元素
	myarrr = append(myarrr, 2)
	// 追加多个元素
	myarrr = append(myarrr, 3, 4)
	// 追加一个切片, ... 表示解包，不能省略
	myarrr = append(myarrr, []int{7, 8}...)
	// 在第一个位置插入元素
	myarrr = append([]int{0}, myarrr...)
	// 在中间插入一个切片(两个元素)
	myarrr = append(myarrr[:5], append([]int{5, 6}, myarrr[5:]...)...)
	fmt.Println(myarr)

}
