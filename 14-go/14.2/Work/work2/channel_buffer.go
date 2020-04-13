package main

import (
	"fmt"
	"time"
)

func main() {
	buf := 10
	ch1 := make(chan int, buf)
	//ch1 := make(chan int)
	go outChan(ch1)
	inChan(ch1)
}

func inChan(ch chan int) {
	fmt.Println("inserting chan")
	ch <- 10
}

func outChan(ch chan int) {
	fmt.Println("waiting for insert")
	time.Sleep(15 * 1e9)
	<-ch
	fmt.Println("output ok")
}

//结果：阻塞和非阻塞的区别，导致输出结果不同了。
