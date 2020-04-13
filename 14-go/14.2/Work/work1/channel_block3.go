package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go outChan(ch1)
	inChan(ch1)
}

func inChan(ch chan int) {
	fmt.Println("inserting chan")
	ch <- 10
}

func outChan(ch chan int) {
	time.Sleep(15 * 1e9)
	fmt.Println("waiting for insert")
	<-ch
	fmt.Println("output ok")
}
