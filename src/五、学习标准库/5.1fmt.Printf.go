/*
1.fmt的三大函数比较

	1.1fmt.Print
	1.2fmt.Println
	1.3fmt.Printf

2.初识fmt.Printf函数

	func Printf(format string, a ...interface{}) (n int, err error) {
	    return Fprintf(os.Stdout, format, a...)
	}

	它的第一个参数是需要格式化的字符串，这个字符串可以是不包含占位符的字符串，也可以是包含站位符的字符串
	占位符：是以%开头的n位短代码，这些短代码根据约定的格式决定着变量输出的格式。
*/
package main

import "fmt"

type demofmtPrint struct{}

func (demo demofmtPrint) demo01() {
	n := 1024
	fmt.Printf("%d的 2 进制： %b \n", n, n)
	fmt.Printf("%d的 8 进制： %o \n", n, n)
	fmt.Printf("%d的 10 进制： %d \n", n, n)
	fmt.Printf("%d的 16 进制： %x \n", n, n)
}

type Profile struct {
	name   string
	gender string
	age    int
}

func (demo demofmtPrint) demo02_normal() {
	var people = Profile{name: "wangbm", gender: "male", age: 27}
	// %v以值的默认格式打印
	fmt.Printf("people: %v\n", people)
	// %T打印值的类型
	fmt.Printf("people: %T\n", people)
	// %#v值的Go语法表示
	fmt.Printf("people: %#v\n", people)
	// %+v类似于%v,但输出结构体时会添加字段名
	fmt.Printf("people: %+v\n", people)
}

func (demo demofmtPrint) demo02_bool() {
	// %t打印布尔值
	fmt.Printf("%t\n", true)
	fmt.Printf("%v\n", false)
}

func (demo demofmtPrint) demo02_string() {
	s := "hello world"
	// %s输出字符串表示(string类型或[]byte)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("%s\n", []byte("hello world"))
	fmt.Printf("%s\n", "hello \\r\\n world")

	// %q双引号围绕的字符串，由Go语法安全地转义
	fmt.Printf("%q \n", "hello world")
	fmt.Printf("%q \n", []byte("hello world"))
	fmt.Printf("%q \n", "hello \r\n world")
}

func (demo demofmtPrint) demo02_pointer() {
	var people = Profile{name: "wbm", gender: "male", age: 27}
	// %p打印指针
	fmt.Printf("people: %p\n", &people)
}

func (demo demofmtPrint) demo02_float() {
	f := 12.34
	// %e 和 %E都是科学计数法表示
	fmt.Printf("f: %e\n", f)
	fmt.Printf("f: %E\n", f)
	// %f有小数部分，但无指数部分
	fmt.Printf("f: %f\n", f)
	// %F等价于%f
	fmt.Printf("f: %F\n", f)
	// %g根据实际情况采用%e或%f格式
	fmt.Printf("f: %g\n", f)
	// %G根据实际情况采用%E或%F格式
	fmt.Printf("f: %G\n", f)
}

func (demo demofmtPrint) demo02_widthAndaccuracy() {
	n := 12.34
	// %w.af  w表示宽度，a表示精度
	// 如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。举例如下：
	fmt.Printf("n: %f\n", n)
	fmt.Printf("n: %9f\n", n)
	fmt.Printf("n: %.2f\n", n)
	fmt.Printf("n: %9.2f\n", n)
	fmt.Printf("n: %9.f\n", n) // output: n:        1
}

func (demo demofmtPrint) demo02_position() {
	var people = Profile{name: "wj", gender: "male", age: 18}
	fmt.Printf("people: %v\n", people)
	// %+v 若值为结构体，则输出将包括结构体的字段名。
	fmt.Printf("people: %+v\n", people)

	// %+q 保证只输出ASCII编码的字符，非 ASCII 字符则以unicode编码表示
}

func (demo demofmtPrint) demo02_alignAndconplete() {
	// 打印的值宽度为5，若不足5个字符，则在前面补空格凑足5个字符
	fmt.Printf("a%5sc\n", "b")
	// 打印的值宽度为5，若不足5个字符，则在后面补空格凑足5个字符
	fmt.Printf("a%-5sc\n", "b")

	// 正负号占位
	fmt.Printf("1% d3\n", 22)
	fmt.Printf("1% d3\n", -22)
}

func main() {
	demo := demofmtPrint{}
	demo.demo02_alignAndconplete()
}
