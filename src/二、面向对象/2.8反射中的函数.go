/*
	1.获取类别
		Kind() Type和Value对象都可以通过kind方法返回对应的接口变量的基础类型ptr struct
	2.进行类型的转换
		Int()
*/

package main

import (
	"fmt"
	"reflect"
)

type Profile struct {
	name string
	age  int
}

func Demo_Kind1() {
	m := Profile{}
	t := reflect.TypeOf(m)
	// fmt.Printf("t: %v\n", t)
	// fmt.Printf("t: %v\n", t.Kind())

	var a interface{}
	t = reflect.TypeOf(a)
	// fmt.Printf("t: %v\n", t)
	// fmt.Printf("t.Kind(): %v\n", t.Kind())
	var s string
	t = reflect.TypeOf(s)
	// fmt.Printf("s: %v\n", t)
	// fmt.Printf("t.Kind(): %v\n", t.Kind())

	var number int
	t = reflect.TypeOf(&number)
	// fmt.Printf("t: %v\n", t)
	// fmt.Printf("t.Kind(): %v\n", t.Kind())
	// fmt.Printf("t.Elem(): %v\n", t.Elem())
	// fmt.Printf("t.Elem().Kind(): %v\n", t.Elem().Kind())

	t = reflect.TypeOf(&a)
	fmt.Printf("t: %v\n", t)
	fmt.Printf("t.Kind(): %v\n", t.Kind())
	fmt.Printf("t.Elem(): %v\n", t.Elem())
	fmt.Printf("t.Elem().Kind(): %v\n", t.Elem().Kind())
}

func Demo_Kind2() {
	m := Profile{}

	t := reflect.TypeOf(&m)
	fmt.Printf("t: %v\n", t)
	fmt.Printf("t.Kind(): %v\n", t.Kind())

	fmt.Printf("t.Elem(): %v\n", t.Elem())
	fmt.Printf("t.Elem().Kind(): %v\n", t.Elem().Kind())
}

func Demo_Int() {
	var a int = 5

	v1 := reflect.ValueOf(a)

	fmt.Printf("v1: %T, %v\n", v1, v1)

	v2 := v1.Int()
	fmt.Printf("v2: %T %v\n", v2, v2)
}

func Demo_Float() {
	var score float64 = 99.5
	v1 := reflect.ValueOf(score)
	fmt.Printf("v1:^%T %v\n", v1, v1)

	v2 := v1.Float()
	fmt.Printf("v2: %T %v\n", v2, v2)
}

func Demo_String() {
	s := "Hello World!"
	v1 := reflect.ValueOf(s)
	fmt.Printf("v1:%T %v\n", v1, v1)
	v2 := v1.String()
	fmt.Printf("v2:%T %v\n", v2, v2)
}

func Demo_Bool() {
	b := true
	v1 := reflect.ValueOf(b)
	fmt.Printf("v1: %T %v\n", v1, v1)

	v2 := v1.Bool()
	fmt.Printf("v2: %T %v\n", v2, v2)
}

func Demo_Pointer() {
	var age int = 23
	v1 := reflect.ValueOf(&age)
	fmt.Printf("v1: %T %v\n", v1, v1)

	v2 := v1.Pointer()
	fmt.Printf("v2: %T %v\n", v2, v2)
}

func Demo_Interface() {
	var age int = 23
	v1 := reflect.ValueOf(age)
	fmt.Printf("v1: %T %v\n", v1, v1)
	v2 := v1.Interface()
	fmt.Printf("v2: %T %v\n", v2, v2)
}

func main() {
	Demo_Interface()
}
