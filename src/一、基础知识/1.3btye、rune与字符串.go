/*
	byte 占用1个字节，就8个比特位，表示范围0~256，本质上和uint8没区别
	rune 占用4个字节，共32个比特位，和int32本质上没有区别
	两者都是字符类型
*/

/*
	byte和rune都是字符类型，多个字符放在一起就组成了字符串，也就是string类型
	string本质上就是一个byte数组
	可以用双引号和反引号表示，反引号包裹的字符串相当于Python中的raw字符串，会忽略里面的转义，所见即所得。
	同时反引号可以不写换行符来表示一个多行的字符串。
*/

package main

import (
	"fmt"
	"unsafe"
)

type demoByte struct {
}

func (demo demoByte) demo01() {
	// byte和uint8本质上没区别
	var a byte = 18
	fmt.Printf("a: %T %v\n", a, a)
	fmt.Printf("a 占用%d 个字节数", unsafe.Sizeof(a))
}

func (demo demoByte) demo02() {
	// rune和uint32本质上没区别
	var a rune = 18
	fmt.Printf("a: %T %v\n", a, a)
	fmt.Printf("unsafe.Sizeof(a): %v\n", unsafe.Sizeof(a))
}

func (demo demoByte) demo03() {
	// byte和rune都是字符类型
	// string本质上就是byte数组
	var a [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("a: %s\n", a)
	var b [5]rune = [5]rune{104, 101, 108, 108, 111}
	fmt.Printf("b: %c\n", b)
}

func (demo demoByte) demo04() {
	// "" 会转义字符串
	// `` 会忽略转义字符
	var mystring01 string = "\\r\\n"
	fmt.Printf("mystring01: %v\n", mystring01)
	var mystring02 string = `\r\n`
	fmt.Printf("mystring02: %v\n", mystring02)
}

func main() {
	demo := demoByte{}
	demo.demo04()

	// var mystr01 string = "hello"
	// var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	// fmt.Printf("mystr01: %s\n", mystr01)
	// fmt.Printf("mystr02: %s\n", mystr02)

	// var mystr01 string = "\\r\\n"
	// var mystr02 string = `\r\n`
	// fmt.Println(mystr01)
	// fmt.Println(mystr02)

}
