/*
1.触发panic
2.捕获 panic
	recover函数(有一个条件，就是它必须在 defer 函数中才能生效，其他作用域下，它是不工作的。)
3.无法跨携程
	但是这个 defer 在多个协程之间是没有效果，在子协程里触发 panic，只能触发自己协程内的 defer，而不能调用 main 协程里的 defer 函数的。


	·panic：抛出异常，使程序崩溃
	·recover：捕获异常，恢复程序或做收尾工作
*/

package main

import "fmt"

func set_data(x int) {
	defer func() {
		// recover() 可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 故意制造数组越界，触发 panic
	var arr [10]int
	arr[x] = 88
}

func main() {
	set_data(20)

	// 如果能执行到这句，说明panic被捕获了
	// 后续的程序能继续运行
	fmt.Println("everything is ok")
}
