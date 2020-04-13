package main

import (
	"fmt"
	"time"
)

func main() {
	stream := pump()
	go suck(stream)
	time.Sleep(1e9)
}
func pump() chan int {
	//通过pump函数返回了一个chan
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}
func suck(ch chan int) {
	//for循环进行读取ch的值
	for {
		fmt.Println(<-ch)
	}
}
