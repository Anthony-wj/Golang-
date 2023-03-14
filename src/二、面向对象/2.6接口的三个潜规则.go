/*
1.对方法的调用限制

	我定义了一个 Phone 的接口，只要求实现 call 方法即可，也就是只要能打电话的设备就是一个电话（好像是一句没用的废话）。
	然后再定义了一个 iPhone 的结构体，该结构体接收两个方法，一个是打电话（ call 函数），一个是发微信（send_wechat 函数）。
	最后一步是关键，我们定义了一个 Phone 接口类型的 phone 对象，该对象的内容是 iPhone 结构体。然后我们调用该对象的 call 方法，一切正常。
	但是当你调用 phone.send_wechat方法，程序会报错，提示我们 Phone 类型的方法没有 send_wechat 的字段或方法。
	那么如何让 phone 可以调用 send_wechat 方法呢？
	答案是可以不显式的声明为 Phone接口类型 ，但要清楚 phone 对象实际上是隐式的实现了 Phone 接口，如此一来，方法的调用就不会受到接口类型的约束。

2.调用接口的隐式转换

	当一个函数接口 interface{} 空接口类型时，我们说它可以接收什么任意类型的参数（江湖上称之为无招胜有招）。
	当你使用这种写法时，Go 会默默地为我们做一件事，就是把传入函数的参数值（注意：Go 语言中的函数调用都是值传递的）的类型隐式的转换成 interface{} 类型。

3.如何进行解耦类型的显式转换

	var a int = 25
	b := interface{}(a)
	如果你想手动对其进行类型转换，可以像下面这样子，就可以将变量 a 的静态类型转换为 interface{} 类型然后赋值给 b
	（此时 a 的静态类型还是 int，而 b 的静态类型为 interface{}）

4.类型断言中的隐式转换

	上面我们知道了，只有静态类型为接口类型的对象才可以进行类型断言。
	而当类型断言完成后，会返回一个静态类型为你断言的类型的对象，也就是说，当我们使用了类型断言，Go 实际上又会默认为我们进行了一次隐式的类型转换。
	验证方法也很简单，使用完一次类型断言后，对返回的对象再一次使用类型断言，Goland 立马就会提示我们新对象 b 不是一个接口类型的对象，不允许进行类型断言。
*/
package main

import "fmt"

type Phone interface {
	call()
}

type iPhone struct {
	name string
}

func (phone iPhone) call() {
	fmt.Println("Hello, iPhone.")
}

func (phone iPhone) send_wechat() {
	fmt.Println("Hello, Wechat.")
}

func printType(i interface{}) {

	switch i.(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}

func main() {
	// var phone Phone
	// phone := iPhone{name: "ming's iphone"}
	// phone.call()
	// phone.send_wechat()
	a := 10
	printType(a)
	switch interface{}(a).(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
	// 当你使用这种写法时，Go 会默默地为我们做一件事，就是把传入函数的参数值（注意：Go 语言中的函数调用都是值传递的）的类型隐式的转换成 interface{} 类型。

}
