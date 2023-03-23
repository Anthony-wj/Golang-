/*
	time.Sleep是一种极不推荐的方式，如何优雅地来处理阻塞问题？
	1.使用信道来标记完成：
		学习了信道后，我们知道，信道可以实现多个协程间的通信，那么我们只要定义一个信道，在任务完成后，往信道中写入true，然后在主协程中获取到true，就认为子协程已经执行完毕。
		单个协程或协程数少的时候使用
	2.sync.WaitGroup

*/

package main

import (
	"fmt"
	"sync"
)

type demoWaitGroup struct{}

func (demo demoWaitGroup) demo01() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	<-done
}
func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d: %d\n", x, i)
	}
}

func (demo demoWaitGroup) demo02() {
	var wg sync.WaitGroup

	wg.Add(2)
	go worker(1, &wg)
	go worker(2, &wg)
	wg.Wait()
}

func main() {
	demo := demoWaitGroup{}
	demo.demo02()
}
