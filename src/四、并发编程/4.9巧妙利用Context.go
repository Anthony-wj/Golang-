/*
	2.为何需要Context
		当一个协程开启后，我们是无法强制关闭它的。
		常见的关闭协程的原因：
			1.goroutine自己跑完结束退出
			2.主进程crash退出，goroutine被迫退出
			3.通过管道发送信号
			第一种属于正常关闭，不在今天讨论范围之内。
			第二种属于异常关闭，应该优化代码。
			第三种才是开发者可以手动控制协程的方法

	3.简单使用Context

	4.根Context

	6.Context注意事项const
		通常 Context 都是做为函数的第一个参数进行传递（规范性做法），并且变量名建议统一叫 ctx

		Context 是线程安全的，可以放心地在多个 goroutine 中使用。

		当你把 Context 传递给多个 goroutine 使用时，只要执行一次 cancel 操作，所有的 goroutine 就可以收到 取消的信号

		不要把原本可以由函数参数来传递的变量，交给 Context 的 Value 来传递。

		当一个函数需要接收一个 Context 时，但是此时你还不知道要传递什么 Context 时，可以先用 context.TODO 来代替，而不要选择传递一个 nil。

		当一个 Context 被 cancel 时，继承自该 Context 的所有 子 Context 都会被 cancel。
*/

package main

import (
	"context"
	"fmt"
	"time"
)

type demoContext struct{}

func (demo demoContext) demo01() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出,停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了,通知监控停止")

	stop <- true
	time.Sleep(5 * time.Second)
}

func monitor(ch chan bool, number int) {
	for {
		select {
		case v := <-ch:
			fmt.Printf("监控器%v,接受到通道值为：%v,监控结束。\n", number, v)
			return
		default:
			fmt.Printf("监控器%v,正在监控中...\n", number)
			time.Sleep(2 * time.Second)
		}
	}
}

func (demo demoContext) demo02() {
	stopSingal := make(chan bool)

	for i := 1; i <= 5; i++ {
		go monitor(stopSingal, i)
	}

	time.Sleep(1 * time.Second)

	close(stopSingal)

	time.Sleep(5 * time.Second)

	fmt.Println("主程序退出")
}

func monitorContext(ctx context.Context, number int) {
	for {
		select {
		//检查ctx.Done()是否刻度，可读就说明context已经取消，你可以清廉goroutine并退出了
		case v := <-ctx.Done():
			fmt.Printf("监控器%v,接收到通道值为：%v,监控结束。\n", number, v)
			return
		default:
			fmt.Printf("监控器%v,正在监控中...\n", number)
			time.Sleep(2 * time.Second)
		}
	}
}

func (demo demoContext) demo03() {
	// 定义一个可取消的context
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 5; i++ {
		go monitorContext(ctx, i)
	}

	time.Sleep(1 * time.Second)
	// 当你想取消context时，只要调用cancel方法即可，这个cancel是我们在创建ctx时返回的第二个值。
	cancel()

	time.Sleep(5 * time.Second)

	fmt.Println("主程序退出...")
}

func main() {
	demo := demoContext{}
	demo.demo03()
}
