package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}
func main() {
	out := make(chan int)
	//先插入
	out <- 2
	//此处输出，已经无值了。
	go f1(out)
	//out <- 2
}
