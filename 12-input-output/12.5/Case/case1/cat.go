package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

//定义了cat函数，当arg为0时，将输入转行为输出
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}
func main() {
	//flag置位
	flag.Parse()
	//如果参数为0，则将输入打印到屏幕。
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
