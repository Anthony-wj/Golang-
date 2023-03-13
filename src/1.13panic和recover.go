/*
1.触发panic
2.捕获 panic
	recover函数(有一个条件，就是它必须在 defer 函数中才能生效，其他作用域下，它是不工作的。)
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
