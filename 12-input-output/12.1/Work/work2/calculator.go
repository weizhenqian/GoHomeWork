/* 编写一个简单的逆波兰式计算器，它接受用户输入的整型数（最大值 999999）和运算符 +、-、*、/。
输入的格式为：number1 ENTER number2 ENTER operator ENTER —> 显示结果
当用户输入字符 ‘q’ 时，程序结束。请使用您在练习11.3中开发的 stack 包。 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"../../../../11-Interface/11.13/Work/stack"
)

var operator string
var quitflag string
var inputReader *bufio.Reader

func main() {
	st1 := new(stack.Stack)
	inputReader = bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Please input Number or operator:")
		input, err := inputReader.ReadString('\n')
		input = input[0 : len(input)-2]
		if err != nil {
			fmt.Println("input error!")
			break
		}
		switch {
		case input == "q":
			fmt.Println("Exit Process!")
			os.Exit(1)
		case input > "0" && input < "999999":
			i, _ := strconv.Atoi(input)
			st1.Push(i)
		case input == "+":
			number1, _ := st1.Pop()
			number2, _ := st1.Pop()
			fmt.Println(number1.(int) + number2.(int))
		}
	}
}

//此代码借鉴了无闻的代码，这个家庭作业的WORK1和WORK2主要区分，就是一个给的值是可以确定的，而另一个是无法确定的。其次，很多os包、strconv用法也得到了锻炼和提升。
