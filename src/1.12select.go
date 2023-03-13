/*
2.避免造成死锁
如果在遍历完所有的 case 后，若没有命中（命中：也许这样描述不太准确，我本意是想说可以执行信道的操作语句）任何一个 case 表达式，就会进入 default 里的代码分支。

但如果你没有写 default 分支，select 就会阻塞，直到有某个 case 可以命中，而如果一直没有命中，select 就会抛出 deadlock 的错误，就像下面这样子。

3.select随机性
之前学过 switch 的时候，知道了 switch 里的 case 是顺序执行的，但在 select 里却不是。

4.select超时
当 case 里的信道始终没有接收到数据时，而且也没有 default 语句时，select 整体就会阻塞，但是有时我们并不希望 select 一直阻塞下去，这时候就可以手动设置一个超时时间。
func makeTimeout(ch chan bool, t int) {
    time.Sleep(time.Second * time.Duration(t))
    ch <- true
}

func main() {
    c1 := make(chan string, 1)
    c2 := make(chan string, 1)
    timeout := make(chan bool, 1)

    go makeTimeout(timeout, 2)

    select {
    case msg1 := <-c1:
        fmt.Println("c1 received: ", msg1)
    case msg2 := <-c2:
        fmt.Println("c2 received: ", msg2)
    case <-timeout:
        fmt.Println("Timeout, exit.")
    }
}
5.读取/写入都可以
上面例子里的 case，好像都只从信道中读取数据，但实际上，select 里的 case 表达式只要求你是对信道的操作即可，不管你是往信道写入数据，还是从信道读出数据。

6.信道关闭也能命中


*/

package main

import (
	"fmt"
)

func main() {
	// c1 := make(chan string, 1)
	// c2 := make(chan string, 1)

	// c2 <- "hello"

	// select {
	// case msg1 := <-c1:
	// 	fmt.Println("c1 received: ", msg1)
	// case msg2 := <-c2:
	// 	fmt.Println("c2 received: ", msg2)
	// default:
	// 	fmt.Println("No data received.")
	// }
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	close(c1)
	for {
		select {
		case <-c1:
			fmt.Println("stop")
			return
		case <-c2:
			fmt.Println("hhh")

		}
	}

}
