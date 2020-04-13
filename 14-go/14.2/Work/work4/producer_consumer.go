//用这种习惯用法写一个程序，有两个协程，第一个提供数字 0，10，20，…90 并将他们放入通道，第二个协程从通道中读取并打印。main() 等待两个协程完成后再结束
package main

import "fmt"

func inChan(ch chan int) {
	count := 10
	for i := 0; i < count; i++ {
		ch <- i * 10
	}
	close(ch)
}

func outChan(ch chan int, done chan bool) {
	for i := range ch {
		fmt.Printf("value:%v\r\n", i)
	}
	done <- true
}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go outChan(ch, done)
	go inChan(ch)
	<-done
}

//写后感：期初，只定义了一组channel，但是无法按照常规逻辑运行。后加了一组channel作为完成后的flag。
