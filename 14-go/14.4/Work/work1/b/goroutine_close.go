/* 在练习 5.4 的 for_loop.go 中，有一个常见的 for 循环打印数字。在函数 tel 中实现一个 for 循环，用协程开始这个函数并在其中给通道发送数字。main() 线程从通道中获取并打印。不要使用 time.Sleep() 来同步： */
package main

import "fmt"

func tel(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i
	}
	//也许你的方案有效，可能会引发运行时的 panic：throw:all goroutines are asleep-deadlock! 为什么会这样？你如何解决这个问题？goroutine_close.go
	defer close(ch)
}

func main() {
	ch1 := make(chan int)
	go tel(ch1)
	for i := range ch1 {
		fmt.Printf("var in ch1 is %v\r\n", i)
	}
}
