/*
Type Assertion（中文名叫：类型断言），通过它可以做到以下几件事情
	检查 i 是否为 nil
	检查 i 存储的值是否为某个类型

第一种：
	t := i.(T)
		这个表达式可以断言一个接口对象（i）里不是 nil，并且接口对象（i）存储的值的类型是 T，如果断言成功，就会返回值给 t，如果断言失败，就会触发 panic。
		如果要断言的接口值是 nil，也会如预期一样会触发panic
第二种：
	t, ok:= i.(T)
		和上面一样，这个表达式也是可以断言一个接口对象（i）里不是 nil，并且接口对象（i）存储的值的类型是 T，如果断言成功，就会返回其值给 t，并且此时 ok 的值 为 true，表示断言成功。
		如果接口值的类型，并不是我们所断言的 T，就会断言失败，但和第一种表达式不同的事，这个不会触发 panic，而是将 ok 的值设为 false ，表示断言失败，此时t 为 T 的零值。
*/

/*
Type Switch
	如果需要区分多种类型，可以使用type switch断言，这个将会比一个一个进行类型断言更简单、直接、高效

*/

package main

import "fmt"

func main() {
	// var i interface{} = 10
	// t1, ok := i.(int)
	// fmt.Println(t1, ok)

	// fmt.Println("=====分隔线=====")

	// t2, ok := i.(string)
	// fmt.Println(t2, ok)
	var a interface{} = "1"
	fmt.Printf("a: %T\n", a)
	switch a.(type) {
	case int:
		fmt.Println("int型")
	case string:
		fmt.Println("string型")
	case bool:
		fmt.Println("bool型")
	default:
		fmt.Println("未知类型")
	}
}
