/*
	拿 switch 后的表达式分别和 case 后的表达式进行对比，只要有一个 case 满足条件，就会执行对应的代码块，然后直接退出 switch - case ，
	如果 一个都没有满足，才会执行 default 的代码块
	1.常用示例
	switch 表达式 {
	    case 表达式1:
	        代码块
	    case 表达式2:
	        代码块
	    default:
	        代码块
	}
	2.一个case多个条件
    month := 2
    switch month {
    case 3, 4, 5:
        fmt.Println("春天")
    case 6, 7, 8:
        fmt.Println("夏天")
    case 9, 10, 11:
        fmt.Println("秋天")
    case 12, 1, 2:
        fmt.Println("冬天")
    default:
        fmt.Println("输入有误...")
    }
	3.case 条件常量不能重复
	当 case 后接的是常量时，该常量只能出现一次。
	4.switch后可接函数
	switch 后面可以接一个函数，只要保证 case 后的值类型与函数的返回值 一致即可。
		// 判断一个同学是否有挂科记录的函数
		// 返回值是布尔类型
		func getResult(args ...int) bool {
			for _, i := range args {
				if i < 60 {
					return false
				}
			}
			return true
		}

		func main() {
			chinese := 80
			english := 50
			math := 100

			switch getResult(chinese, english, math) {
			// case 后也必须 是布尔类型
			case true:
				fmt.Println("该同学所有成绩都合格")
			case false:
				fmt.Println("该同学有挂科记录")
			}
		}
	5.switch可不接表达式
	switch 后可以不接任何变量、表达式、函数。
	当不接任何东西时，switch - case 就相当于 if - elseif - else
		score := 30
		switch {
			case score >= 95 && score <= 100:
				fmt.Println("优秀")
			case score >= 80:
				fmt.Println("良好")
			case score >= 60:
				fmt.Println("合格")
			case score >= 0:
				fmt.Println("不合格")
			default:
				fmt.Println("输入有误...")
		}
	6.switch的穿透能力
	关键字fallthrough，
		s := "hello"
		switch {
		case s == "hello":
			fmt.Println("hello")
			fallthrough
		case s != "world":
			fmt.Println("world")
	需要注意的是，fallthrough 只能穿透一层，意思是它让你直接执行下一个case的语句，而且不需要判断条件。
}
*/
/*
package main

import "fmt"

func main() {
	education := "本科"

	switch education {
	case "博士":
		fmt.Println("我是博士")
	case "研究生":
		fmt.Println("我是研究生")
	case "本科":
		fmt.Println("我是本科生")
	case "大专":
		fmt.Println("我是大专生")
	case "高中":
		fmt.Println("我是高中生")
	default:
		fmt.Println("学历未达标..")
	}
}
*/
package main

import "fmt"

type demoSwitch struct {
}

func (demo demoSwitch) demo01() {
	education := "本科"

	switch education {
	case "博士":
		fmt.Println("我是博士生")
	case "研究生":
		fmt.Println("我是研究生")
	case "本科":
		fmt.Println("我是本科生")
	case "小学":
		fmt.Println("我是小学生")
	}
}

func (demo demoSwitch) demo02() {
	// case后接多个示例
	month := 2
	switch month {
	case 3, 4, 5:
		fmt.Println("春天")
	case 6, 7, 8:
		fmt.Println("夏天")
	case 9, 10, 11:
		fmt.Println("秋天")
	case 12, 1, 2:
		fmt.Println("冬天")
	}
}

func main() {
	demo := demoSwitch{}
	demo.demo01()
}
