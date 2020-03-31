//扩展 cat.go 例子，使用 flag 添加一个选项，目的是为每一行头部加入一个行号。使用 cat -n test 测试输出
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

//定义一个默认值为false的位置参数
var num *bool

//定义了cat函数，当arg为0时，将输入转行为输出
func cat(r *bufio.Reader) {
	if *num {
		i := 0
		for {
			i += 1
			buf, err := r.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stdout, "%v%s", i, buf)
		}
		return
	} else {
		for {
			buf, err := r.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
		return
	}

}
func main() {
	num = flag.Bool("n", false, "add line number")
	//flag置位
	flag.Parse()

	//如果参数为0，则将输入打印到屏幕。
	fmt.Println(flag.NArg())
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	//如果参数不为0，则打开文件，将文件内容输出。
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
