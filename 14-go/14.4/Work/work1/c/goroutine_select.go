//解决方案2：使用一个额外的通道传递给协程，然后在结束的时候随便放点什么进去。main() 线程检查是否有数据发送给了这个通道，如果有就停止：goroutine_select.go
package main

import (
	"fmt"
	"os"
)

func tel(ch1 chan int, ch2 chan bool) {
	for i := 0; i < 20; i++ {
		ch1 <- i
	}
	defer func() { ch2 <- true }()
	//defer close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool)
	go tel(ch1, ch2)
	//for循环为了保证select能够完整运行
	for {
		select {
		case v := <-ch1:
			fmt.Printf("var in ch1 is %v", v)
		case _ = <-ch2:
			os.Exit(0)
		}
	}
}
