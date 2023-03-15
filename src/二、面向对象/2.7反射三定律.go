/*
	2两种类型：Type和Value
		2.1reflect包中的两种类型
			reflect.Type 从源码上来看，是以一个接口的形式存在的
			reflect.Value 从源码上来看，是以一个结构体的形式存在的
		2.2
			我们知道了一个接口变量，实际上都是由一 pair 对（type 和 data）组合而成，pair 对中记录着实际变量的值和类型。也就是说在真实世界里，type 和 value 是合并在一起组成 接口变量的。
			而在反射的世界里，type 和 data 却是分开的，他们分别由 reflect.Type 和 reflect.Value 来表现。
	3反射三定律：
		第一定律：反射可以将接口类型变量转换为"反射类型对象"
			为了实现从接口变量到反射对象的转换，需要提高reflect包里很重要的两个方法：
				1.reflect.TypeOf(i)
				2.reflect.ValueOf(i)
				由于 TypeOf 和 ValueOf 两个函数接收的是 interface{} 空接口类型，而 Go 语言函数都是值传递，因此Go语言会将我们的类型隐式地转换成接口类型。
		第二定律：反射可以将"反射类型对象"转换为接口类型变量
			i := v.Interface()
			i的静态类型是interface{}，如果要转成原始类型，需要再加个断言转换一下i.(int)
		第三定律：传递不是变量的指针，在函数内部的修改是不会影响到原始变量的
			因此在反射的规则里：
				不是接收变量指针创建的反射对象，是不具备『可写性』的
				是否具备『可写性』，可使用 CanSet() 来获取得知
				对不具备『可写性』的对象进行修改，是没有意义的，也认为是不合法的，因此会报错
			要让反射对象具备可写性，需要注意两点
				1.创建反射对象时传入变量的指针
				2.使用Elem()函数返回指针指向的数据
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var age int = 23
	// fmt.Printf("原始接口变量的类型为 %T,age: %v\n", age, age)

	// t := reflect.TypeOf(age)
	// v := reflect.ValueOf(age)
	// fmt.Printf("T: %T\n", t)
	// fmt.Printf("T: %T\n", v)
	// i := v.Interface()
	// fmt.Printf("i:%T %v\n", i, i)
	/*
		如此我们完成了从接口类型变量到反射对象的转换。
		等等，上面我们定义的 age 不是 int 类型的吗？第一法则里怎么会说是接口类型的呢？
		关于这点，其实在上一节（关于接口的三个 『潜规则』）已经提到过了，由于 TypeOf 和 ValueOf 两个函数接收的是 interface{} 空接口类型，而 Go 语言函数都是值传递，因此Go语言会将我们的类型隐式地转换成接口类型。
	*/
	var name string = "Go编程时光"

	v1 := reflect.ValueOf(&name)

	fmt.Printf("v1.CanSet(): %v\n", v1.CanSet())
	v2 := v1.Elem()
	fmt.Printf("v2.CanSet(): %v\n", v2.CanSet())
	v2.SetString("Python编程时光")
	fmt.Printf("v2: %v\n", v2)

}
