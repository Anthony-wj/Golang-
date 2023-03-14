/*
		之前学过的数据类型中，数组与切片只能存储同一类型的变量。
		若要存储多个类型的变量，就需要用到结构体，它是将多个任意类型的变量组合在一起的数据类型。

		每个变量都成为该结构体的成员变量。

		可以理解为 Go语言 的结构体struct和其他语言的class有相等的地位，但是Go语言放弃大量面向对象的特性，所有的Go语言类型除了指针类型外，都可以有自己的方法,提高了可扩展性。

		在 Go 语言中没有没有 class 类的概念，只有 struct 结构体的概念，因此也没有继承，本篇文章，带你学习一下结构体相关的内容。

		1.定义结构体
		type 结构体名 struct {
	    属性名   属性类型
	    属性名   属性类型
	    ...
		}
		规则一：
			当最后一个字段和结果不在同一行时，, 不可省略。
			反之，不在同一行，就可以省略。
		规则二：字段名要嘛全写，要嘛全不写，不能有的写，有的不写。
		规则三；初始化结构体，并不一定要所有字段都赋值，未被赋值的字段，会自动赋值为其类型的零值。
		2.绑定方法
		func (person Profile) FmtProfile() {
			fmt.Printf("名字：%s\n", person.name)
			fmt.Printf("年龄：%d\n", person.age)
			fmt.Printf("性别：%s\n", person.gender)
		}
		3.参数传递方式
			当你想要在方法内改变实例的属性的时候，必须使用指针做为方法的接收者。
			func (person *Profile) increase_age() {
				person.age += 1
			}

至此，我们知道了两种定义方法的方式：

	以值作为方法接收者
	以指针作为方法接收者

		4.结构体继承
			在 Go 语言中，把一个结构体嵌入到另一个结构体的方法，称之为组合。

		5.内部方法与外部方法
			在 Go 语言中，函数名的首字母大小写非常重要，它被来实现控制对方法的访问权限。
				当方法的首字母为大写时，这个方法对于所有包都是Public，其他包可以随意调用
				当方法的首字母为小写时，这个方法是Private，其他包是无法访问的。
		6.三种实例化方法
			一、正常实例化
			xm := Profile{
				name: "小明",
				age: 18,
				gender: "male",
			}
			二、使用new
			xm := new(Profile) // 等价于: var xm *Profile = new(Profile)
			三、使用&
			var xm *Profile = &Profile{}
		7.选择器的妙用
			从一个结构体实例对象中获取字段的值，通常都是使用 . 这个操作符，该操作符叫做 选择器。
			当你对象是结构体对象的指针时，你想要获取字段属性时，按照常规理解应该这么做
			但还有一个更简洁的做法，可以直接省去 * 取值的操作，选择器 . 会直接解引用，示例如下
*/
package main

import "fmt"

type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile
	father *Profile
}

func (person Profile) FmtProfile() {
	fmt.Printf("名字: %s\n", person.name)
	fmt.Printf("年龄: %d\n", person.age)
	fmt.Printf("性别: %s\n", person.gender)
}

func (person *Profile) increase_age() {
	person.age += 1
}

type company struct {
	companyName string
	companyAddr string
}
type staff struct {
	name     string
	age      int
	gender   string
	position string
	company
}

func main() {
	// myself := Profile{name: "小明", age: 18, gender: "male"}
	// myself.increase_age()
	// myself.FmtProfile()
	// myCom := company{
	// 	companyName: "Tencent",
	// 	companyAddr: "深圳市南山区",
	// }
	// staffInfo := staff{
	// 	name:     "小明",
	// 	age:      28,
	// 	gender:   "男",
	// 	position: "云计算开发工程师",
	// 	company:  myCom,
	// }
	// fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.companyName)
	// fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.company.companyName)
	// xm := Profile{
	// 	name:   "小明",
	// 	age:    18,
	// 	gender: "male",
	// }
	xm := new(Profile)
	xm.name = "peanuts"
	xm.FmtProfile()

}
