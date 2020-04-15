/* 从示例 6.10 的斐波那契程序开始，制定解决方案，使斐波那契周期计算独立到协程中，并可以把结果发送给通道。

结束的时候关闭通道。main() 函数读取通道并打印结果：goFibonacci.go

使用练习 6.9 中的算法写一个更短的 gofibonacci2.go

使用 select 语句来写，并让通道退出（gofibonacci_select.go）

注意：当给结果计时并和 6.10 对比时，我们发现使用通道通信的性能开销有轻微削减；这个例子中的算法使用协程并非性能最好的选择；但是 gofibonacci3 方案使用了 2 个协程带来了 3 倍的提速。 */
package main

import (
	"fmt"
)

func main() {
	var ch1 chan int
	//此处没有make创建，也没有报错，但是程序执行就是有问题。调试了很久，对比了之前代码才发现这个问题。
	ch1 = make(chan int)
	count := 10
	go fibonacciChan(count, ch1)
	for v := range ch1 {
		fmt.Println(v)
	}
	/* 	for {
		if v, ok := <-ch1; ok {
			fmt.Printf("v:%v", v)
		} else {
			fmt.Printf("end")
			os.Exit(0)
		}
	} */
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func fibonacciChan(count int, ch chan int) {
	for i := 0; i <= count; i++ {
		res := fibonacci(i)
		ch <- res
	}
	defer close(ch)
}
