/*
	Go编译器，对于 { 和 } 的位置有严格的要求，它要求 else if （或 else）和 两边的花括号，必须在同一行。

	由于 Go是 强类型，所以要求你条件表达式必须严格返回布尔型的数据（nil 和 0 和 1 都不行，具体可查看《详解数据类型：字典与布尔类型》）。
	1.单分支：
    age := 20
    if age > 18 {
        fmt.Println("已经成年了")
    }
	如果条件里需要满足多个条件，可以使用 && 和 ||
		&&：表示且，左右都需要为true，最终结果才能为 true，否则为 false
		||：表示或，左右只要有一个为true，最终结果即为true，否则 为 false
	2.多分支 if else if else
	age := 20
    if age > 18 {
        fmt.Println("已经成年了")
    } else if age >12 {
        fmt.Println("已经是青少年了")
    } else {
        fmt.Println("还不是青少年")
    }
	3.高级写法
	在 if 里可以允许先运行一个表达式，取得变量后，再对其进行判断，比如第一个例子里代码也可以写成这样
	if age := 20;age > 18 {
        fmt.Println("已经成年了")
    }
*/

package main

import "fmt"

type demoIf struct{}

func (demo demoIf) demo01() {
	a := 59
	if a < 60 {
		fmt.Println("不及格")
	} else {
		fmt.Println("及格")
	}
}

func (demo demoIf) demo02() {
	a := 85
	if a < 60 {
		fmt.Println("不及格")
	} else if a < 80 {
		fmt.Println("良好")
	} else {
		fmt.Println("优秀")
	}
}

func (demo demoIf) demo03() {
	a := 65
	if a < 60 {
		fmt.Println("挂科了")
	} else {
		if a < 90 {
			fmt.Println("绩点低喽")
		} else {
			fmt.Println("有保研希望")
		}
	}
}

func main() {
	demo := demoIf{}
	demo.demo03()
}
