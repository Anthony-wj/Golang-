/*
	在面向对象的领域里，接口一般这样定义：接口定义一个对象的行为。接口只指定了对象应该做什么，至于如何实现这个行为（即实现细节），则由对象本身去确定。

	在 Go 语言中，接口就是方法签名（Method Signature）的集合。当一个类型定义了接口中的所有方法，我们称它实现了该接口。这与面向对象编程（OOP）的说法很类似。接口指定了一个类型应该具有的方法，并由该类型决定如何实现这些方法。
*/

/*
1.如何定义接口

	使用type关键字来定义接口

	type Phone interface {
		call()
	}

2.如何实现接口

	如果有一个类型/结构体，实现了一个接口要求的所有方法，这里 Phone 接口只有 call方法，所以只要实现了 call 方法，我们就可以称它实现了 Phone 接口。

	意思是如果有一台机器，可以给别人打电话，那么我们就可以把它叫做电话。

	这个接口的实现是隐式的，不像 JAVA 中要用 implements 显示说明。

	继续上面电话的例子，我们先定义一个 Nokia 的结构体，而它实现了 call 的方法，所以它也是一台电话。
*/
package main

import (
	"fmt"
	"strconv"
)

// 定义一个接口
type Good interface {
	settleAccount() int
	orderInfo() string
}

type Phone struct {
	name     string
	quantity int
	price    int
}

func (phone Phone) settleAccount() int {
	return phone.quantity * phone.price
}
func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

type FreeGift struct {
	name     string
	quantity int
	price    int
}

func (gift FreeGift) settleAccount() int {
	return 0
}
func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}
func main() {
	iPhone := Phone{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}
	earphones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}

	goods := []Good{iPhone, earphones}
	fmt.Printf("goods: %v\n", goods)
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单总共需要支付 %d 元", allPrice)
}
