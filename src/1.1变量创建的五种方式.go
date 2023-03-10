package main

func main() {
	// 1 一行声明一个变量 var <name> type
	// var age int
	// var name = "wujun" // 不需要声明变量的类型，编译器会推到
	// var score float32 = 0.89

	// 2 多个变量一起声明 var ()
	// var (
	// 	name string
	// 	age int
	// 	address string
	// )

	// 3 声明并初始化变量 :=   -->至适用于函数内部
	// name := "wujun"   // 等价与 var name string = "wujun" 和 var name = "wujun"

	// 4 声明并初始化多个变量  ->常用于变量的交换
	// name, age := "wujun", 18
	// var a = 10
	// var b = 20
	// a, b = b, a

	// 5 new函数声明一个指针变量  -->new(Type) 初始化为Type类型的零值，然后返回变量地址，返回的指针类型为*Type
	// name := new(string)
	// fmt.Printf("name: %v\n", name)
	// fmt.Printf("*name: %v\n", *name)

	// plus 匿名变量 _ --》也称为占位符
	// 通常我们接受必须接受但用不到的值
	// 优点：1、不分配内存，不占用内存空间。 2、不需要为变量名命名困扰。 3、多次声明不会有任何问题。
}
