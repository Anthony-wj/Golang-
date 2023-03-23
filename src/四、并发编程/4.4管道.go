/*
	如果说goroutine是Go语言程序的并发体的话，namechannel是他们之间的通信机制。
	信道就是一个管道，连接多个goroutine程序，它是一种队列式的数据解耦股，遵循先入先出的规则。
	1.信道的定义与使用
		c := make(chan int)
	2.信道的容量与长度
		一般创建信道都是使用make函数，make函数接受两个参数
		· 第一个参数：必填，指定信道类型
		· 第二个参数：选填，默认为0，指定信道的容量(可缓存多少数据)

		对于信道容量：
			当容量为0时，说明信道不能存放数据，在发送数据时，必须立马有人接收，否则会报错，此时的信道称之为无缓冲信道。
			当容量为1时，说明信道只能缓存一个数据，若信道中已有一个数据，此时再往里发送数据，会造成程序阻塞，利用这一点可以利用信道来做锁。
			当容量大于1时，信道中可以存放多个数据，可以用于多个协程之间的通信管道，共享资源。

		信道的容量cap，信道的长度len
	3.缓冲信道与无缓冲信道
		按照是否缓冲数据可分为：缓冲信道 与 无缓冲信道。
		缓冲信道：
			允许信道里存储一个或多个数据，这意味着，这设置了缓冲区后，发送端和接收端可以处于异步的状态。
		无缓冲信道：
			发送端和接收端是同步运行的
	4.双向信道与单向信道
		双向信道：
			默认情况下定义的信道都是双向的
		单向信道：


*/

package main

import (
	"fmt"
	"time"
)

type demoChannel struct{}

func (demo demoChannel) demo01() {
	// channel的定义：方法1
	var c chan int // 声明后的信道，其零值是nil，无法直接使用，必须配合make函数进行初始化
	c = make(chan int)
	fmt.Printf("c: %v\n", c)
}

func (demo demoChannel) demo02() {
	// channel的定义：方法2
	c := make(chan int)
	fmt.Printf("c: %v\n", c)
}

func (demo demoChannel) demo03() {
	pipline := make(chan int, 1)
	pipline <- 200
	mydate := <-pipline
	fmt.Printf("mydate: %v\n", mydate)
	close(pipline)
	x, ok := <-pipline
	if ok {
		fmt.Printf("x: %v\n", x)
	} else {
		fmt.Printf("ok: %v\n", ok)
	}
}

func (demo demoChannel) demo04() {
	pipline := make(chan int, 10)
	pipline <- 100
	fmt.Printf("cap(pipline): %v\n", cap(pipline))
	fmt.Printf("len(pipline): %v\n", len(pipline))
}

func (demo demoChannel) demo05() {
	// 定义只读信道
	var pipline = make(chan int)
	type Receiver = <-chan int
	// var receiver Receiver = pipline

	// 定义只写信道
	type Sender = chan<- int
	// var sender Sender = pipline
	go func() {
		var sender Sender = pipline
		for i := 0; i < 100; i++ {
			fmt.Printf("准备发送数据: %v\n", i)
			sender <- i
		}
	}()

	go func() {
		var receiver Receiver = pipline
		for {
			num := <-receiver
			fmt.Printf("接收到的数据是: %d", num)
		}
	}()
}

func main() {
	demo := demoChannel{}
	demo.demo05()
	time.Sleep(time.Second)
}
