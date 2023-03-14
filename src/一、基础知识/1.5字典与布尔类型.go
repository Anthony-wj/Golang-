/*
		字典(Map类型)，是由若干个key:value这样的键值对映射组合在一起的数据结构。
		它是哈希表的一个实现，这就要求它的每个映射里的key，都是唯一的，可以用== 和 != 来进行判断，换句话来说，key必须是可哈希的。
	    什么叫可哈希？简单来说，就是一个不可变对象，都可以用一个哈希值来唯一表示，这样的不可变对象比如字符串(可以说除了切片、字典、函数之外的其他内建类型都算)
		声明：map[KEY_TYPE]VALUE_TYPE
*/

/*
		添加元素 scores["Math"] = 99
		替换元素 scores["Math"] = 59
		读取元素 scores["Math"]
		删除元素 delete(scores, "Math")
	    当访问一个不存在的key时，并不会直接报错，而是会返回这个value的零值，如果value是int型，则返回0
*/

/*
布尔类型：
关于布尔值，无非就两个值：true 和 false
在Go中，真值用true表示，不但不与1相等，并且更加严格，不同类型无法进行比较，而价值用false表示，同样与0无法比较。
在Go中，对逻辑值取反使用!符号，在Python中使用not符号
*/
package main

import "fmt"

func main() {
	// 声明1方法  -->不建议拆开
	// var scores map[string]int = map[string]int{"English": 80, "Chinese": 59}
	// 声明2方法
	// scores := map[string]int{"English": 90, "Math": 60}
	// 声明3方法  --》推荐
	scores := make(map[string]int)
	scores["English"] = 85
	scores["Chinese"] = 59

	// 判断key是否存在
	// math, ok := scores["Math"]
	// if ok {
	// 	fmt.Printf("math: %v\n", math)
	// } else {
	// 	fmt.Println("math不存在")
	// }

	// if math, ok := scores["math"]; ok {
	// 	fmt.Printf("math 的值是: %d", math)
	// } else {
	// 	fmt.Println("math 不存在")
	// }

	// 遍历字典
	// 1获取key 和 value
	for subject, score := range scores {
		fmt.Printf("key: %s, value: %d\n", subject, score)
	}
	// 2只获取key
	for subject := range scores {
		fmt.Printf("subject: %v\n", subject)
	}
	// 3.只获取value
	for _, score := range scores {
		fmt.Printf("score: %v\n", score)

	}
}
