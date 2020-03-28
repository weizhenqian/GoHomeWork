/* 编写一个程序，从键盘读取输入。当用户输入 ‘S’ 的时候表示输入结束，这时程序输出 3 个数字：
i) 输入的字符的个数，包括空格，但不包括 ‘\r’ 和 ‘\n’
ii) 输入的单词的个数
iii) 输入的行数 */
//第一版本采用fmt的Scanln实现
package main

import (
	"fmt"
	"strings"
)

var (
	inputlist []string
	input     string
)

func main() {
	for true {
		fmt.Println("Please input some bytes:")
		fmt.Scanln(&input)
		inputlist = append(inputlist, input)
		if input == "S" {
			break
		}
	}
	result1 := strings.Join(inputlist, "")
	fmt.Println(len(result1))
	fmt.Println(len(inputlist))
}
