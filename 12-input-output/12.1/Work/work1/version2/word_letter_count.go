/* 编写一个程序，从键盘读取输入。当用户输入 ‘S’ 的时候表示输入结束，这时程序输出 3 个数字：
i) 输入的字符的个数，包括空格，但不包括 ‘\r’ 和 ‘\n’
ii) 输入的单词的个数
iii) 输入的行数 */
//第一版本采用fmt的Scanln实现
package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("Please enter some input:")
		input, err = inputReader.ReadString('\n')
		//此处需要对input的值进行处理，过滤掉windows上的\r\n字符，以正确匹配
		input = input[0 : len(input)-2]
		if input == "S" {
			break
		} else {
			fmt.Printf("The input was:%s\n", input)
		}
	}
}
