//开启一个协程来计算2个整数的合并等待计算结果并打印出来。
package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int)
	go func(x, y int) {
		res := x + y
		fmt.Printf("result is %v", res)
		ch <- 1
	}(1, 2)
	<-ch
}
