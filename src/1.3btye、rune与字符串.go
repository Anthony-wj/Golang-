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
)

func main() {
	// var a byte = 65
	// fmt.Printf("a: %c\n", a)
	// var b uint = 66
	// fmt.Printf("b: %c\n", b)
	// var c rune = 'A'
	// fmt.Printf("a 占用%d 个字节数, c 占用%d 个字节数", unsafe.Sizeof(a), unsafe.Sizeof(c))

	// var mystr01 string = "hello"
	// var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	// fmt.Printf("mystr01: %s\n", mystr01)
	// fmt.Printf("mystr02: %s\n", mystr02)

	var mystr01 string = "\\r\\n"
	var mystr02 string = `\r\n`
	fmt.Println(mystr01)
	fmt.Println(mystr02)

}
