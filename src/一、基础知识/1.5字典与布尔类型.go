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

type demoMap struct {
}

func (demo demoMap) demo01() {
	var scores map[string]int
	if scores == nil {
		scores = make(map[string]int)
	}
	scores["math"] = 90
	fmt.Printf("scores: %v\n", scores)
}

func (demo demoMap) demo02() {
	scores := map[string]int{"math": 60}
	fmt.Printf("scores: %v\n", scores)
}

func (demo demoMap) demo03() {
	// 强力推荐
	scores := make(map[string]int)
	scores["math"] = 91
	fmt.Printf("scores: %v\n", scores)
}
func (demo demoMap) demo04() {
	// 判断key是否存在
	scores := map[string]int{"math": 66}
	if math, ok := scores["matth"]; ok {
		fmt.Printf("math: %v\n", math)
	} else {
		fmt.Println("math不存在")
	}
}

func (demo demoMap) demo05() {
	// 遍历字典
	scores := make(map[string]int)
	scores["Math"] = 100
	scores["English"] = 99
	scores["Chinese"] = 98
	// k v遍历
	for k, v := range scores {
		fmt.Printf("k: %v\t", k)
		fmt.Printf("v: %v\n", v)
	}

	// 只遍历key
	for k := range scores {
		fmt.Printf("k: %v\n", k)
	}

	// 只遍历value
	for _, v := range scores {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	demo := demoMap{}
	demo.demo05()
}
