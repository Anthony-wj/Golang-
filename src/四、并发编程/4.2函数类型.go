/*
	函数类型是一种很特殊的类型，他表示所有拥有同样的入参类型和返回值类型的函数集合
*/

package main

import "fmt"

type Greeting func(name string) string

func (g Greeting) say(n string) {
	fmt.Println(n)
}

func english(name string) string {
	return "Hello, " + name
}

func main() {
	greet := Greeting(english)
	greet.say("helo")
}
