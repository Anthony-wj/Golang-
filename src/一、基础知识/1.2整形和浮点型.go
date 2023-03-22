/*
	1.整形
	分为int和uint，区别在于u表示无符号。
	int8和uint8为例，8代表8bit，能够表示的数值个数都为2^8=256。uint8表示0~255，int8表示-128~127
	int并没有指定它的位数，在32位和64位机器中不同
		在32位系统下，int和uint都占用4个字节，也就是32位。
		在64位系统下，int和uint都占用8个字节，也就是64位。
	所以，在有些场景中，应当避免使用int和uint，转而使用更加精确地int32和int64。

	2.浮点型
	float32和float64
*/

package main

import "fmt"

type demoInt struct {
}

func (demo demoInt) demo01() {
	var a int = 18
	fmt.Printf("a: %T %v\n", a, a)
}

func (demo demoInt) demo02() {
	var a int8 = -128 // 最小为-128，最大为127
	fmt.Printf("a: %T %v\n", a, a)
}

func (demo demoInt) demo03() {
	var a uint8 = 255 //最小为0， 最大为255
	fmt.Printf("a: %T %v\n", a, a)
}

func (demo demoInt) demo04() {
	var a int64 = 1844674407370955161

	fmt.Printf("a: %v\n", a)
}

func (demo demoInt) demo05() {
	var a float32 = 0.000000000000000000000000000000000001
	fmt.Printf("a: %v\n", a)
}

func main() {
	demo := demoInt{}
	demo.demo05()
}
