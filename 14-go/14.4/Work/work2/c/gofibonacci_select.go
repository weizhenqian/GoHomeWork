package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool)
	count := 10
	go fibonacciChan(count, ch1, ch2)
	for {
		select {
		case v := <-ch1:
			fmt.Printf("v:%v\r\n", v)
		case <-ch2:
			//close(ch2) Panic:close of closed channel
			return
		}
	}
}
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func fibonacciChan(count int, ch1 chan int, ch2 chan bool) {
	for i := 0; i <= count; i++ {
		res := fibonacci(i)
		ch1 <- res
	}
	ch2 <- true
}
