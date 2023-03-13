/*
	for [condition |  ( init; condition; increment ) | Range]
	{
	statement(s);
	}
	可以看到 for 后面，可以接三种类型的表达式。
	1.接一个条件表达式
	2.接三个表达式
	3.接一个 range 表达式
	4.不接表达式(在 Go 语言中，没有 while 循环，如果要实现无限循环，也完全可以 for 来实现。
		当你不加任何的判断条件时， 就相当于你每次的判断都为 true，程序就会一直处于运行状态，但是一般我们并不会让程序处于死循环，
		在满足一定的条件下，可以使用关键字 break 退出循环体，也可以使用 continue 直接跳到下一循环。)
*/

/*
	for range
	 range后可接数组、切片、字符串等。
	 由于range会返回两个值：索引和数据，若你后面的代码用不到索引，需要使用_表示。

*/

package main

import "fmt"

func main() {
	// a := 1
	// for a < 5 {
	// 	fmt.Printf("a: %v\n", a)
	// 	a++
	// }
	// for i := 1; i < 5; i++ {
	// 	fmt.Printf("i: %v\n", i)
	// }
	// s := [...]string{"world", "python", "go"}
	// for i, data := range s {
	// 	fmt.Printf("i: %v ", i)
	// 	fmt.Printf("data: %s\n", data)
	// }

	i := 1
flag:
	if i < 5 {
		fmt.Printf("i: %v\n", i)
		i++
		goto flag
	}
	i = 1
	for {
		if i > 5 {
			goto flag1
		}
		fmt.Printf("i: %v\n", i)
		i++
	}
flag1:
	for i <= 20 {
		if i%2 == 1 {
			i++
			goto flag1
		}
		fmt.Printf("i: %v\n", i)
		i++
	}
}
