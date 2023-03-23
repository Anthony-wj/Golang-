/*
	1.互斥锁Mutex
		v1:var lock *sync.Mutex
		lock = new(sync.Mutex)

		v2:lock := &sync.Mutex{}
		同一协程里，不要在尚未解锁时再次使加锁
		同一协程里，不要对已解锁的锁再次解锁
		加了锁后，别忘了解锁，必要时使用 defer 语句

	2.读写锁RWMutex
		Mutex是一个傻瓜式的锁，RWMutex将程序对资源的访问分为读操作和写操作
			· 为了保证数据的安全，它规定了当有人还在读取数据（即读锁占用）时，不允计有人更新这个数据（即写锁会阻塞）
			· 为了保证程序的效率，多个人（线程）读取数据（拥有读锁）时，互不影响不会造成阻塞，它不会像 Mutex 那样只允许有一个人（线程）读取同一个数据。
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type demoMutex struct{}

func add(count *int, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		*count += 1
	}
	wg.Done()
}

func (demo demoMutex) demo01() {
	var wg sync.WaitGroup
	count := 0
	wg.Add(3)
	go add(&count, &wg)
	go add(&count, &wg)
	go add(&count, &wg)
	wg.Wait()
	fmt.Printf("count: %v\n", count)
}

func add2(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*count += 1
		lock.Unlock()
	}
	wg.Done()
}

func (demo demoMutex) demo02() {
	var wg sync.WaitGroup
	lock := &sync.Mutex{}
	count := 0
	wg.Add(3)
	go add2(&count, &wg, lock)
	go add2(&count, &wg, lock)
	go add2(&count, &wg, lock)
	wg.Wait()
	fmt.Printf("count: %v\n", count)
}

func (demo demoMutex) demo03() {
	lock := &sync.RWMutex{}
	lock.Lock()

	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("第 %v 个协程准备开始....\n", i)
			lock.RLock()
			fmt.Printf("第 %v 个协程获得读锁, sleep 1s 后释放\n", i)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("准备释放写锁, 读锁不再阻塞")
	lock.Unlock()

	lock.Lock()
	fmt.Println("程序退出...")
	lock.Unlock()
}

func main() {
	demo := demoMutex{}
	demo.demo03()
}
