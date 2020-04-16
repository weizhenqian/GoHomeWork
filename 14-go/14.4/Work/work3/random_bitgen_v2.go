//此脚本更加简洁，通过select case 进行0/1的随机赋值
package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go func() {
		//此处其实无需i:=range这么麻烦，因为生产者是死循环，所以消费者也可以设置为死循环
		/* 		for i := range ch1 {
			fmt.Println(i)
		} */
		for {
			fmt.Println(<-ch1)
		}
	}()

	for {
		//select为什么能够保证0/1进行随机的赋值，这个让人很好奇？
		select {
		case ch1 <- 0:
		case ch1 <- 1:
		default:
			ch1 <- 2
		}
	}
}
