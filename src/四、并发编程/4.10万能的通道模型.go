/*
	如何优雅地关闭channel
		1.有隐患且不优雅地做法
			使用recover(),包装个safeClose函数
*/

package main

func SafeClose(ch chan bool) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()
	close(ch)

	return true
}

func SafeSend(ch chan bool, value bool) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()

	ch <- value  // 如果ch已关闭，则产生一个恐慌。
	return false // <=> closed = false; return
}

func main() {
	c := make(chan bool)
	go func() {
		for i := 0; i < 25; i++ {
			SafeSend(c, true)
		}
	}()
	SafeClose(c)
}
