/* 在函数式编程语言中，一个 map-function 是指能够接受一个函数原型和一个列表，并使用列表中的值依次执行函数原型，公式为：map ( F(), (e1,e2, . . . ,en) ) = ( F(e1), F(e2), ... F(en) )。

编写一个函数 mapFunc 要求接受以下 2 个参数：

一个将整数乘以 10 的函数
一个整数列表
最后返回保存运行结果的整数列表。 */
package main

import "fmt"

func upTen(a int) int {
	return a * 10
}

func mapFunc(F func(int) int, ia []int) []int {
	lia := len(ia)
	for i := 0; i < lia; i++ {
		ia[i] = F(ia[i])
	}
	return ia
}

func main() {
	intarray := []int{1, 2, 3, 4, 5}
	mapFunc(upTen, intarray)
	fmt.Println(intarray)
}
