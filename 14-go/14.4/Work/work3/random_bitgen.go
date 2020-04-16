//做一个随机位生成器，程序可以提供无限的随机 0 或者 1 的序列：random_bitgen.go
package main

import (
	"fmt"
	"math/rand"
)

func numCreate(ch chan int) {
	for {
		ch <- rand.Intn(2)
	}
}

func main() {
	ch1 := make(chan int)
	go numCreate(ch1)
	for i := range ch1 {
		fmt.Println(i)
	}
}
