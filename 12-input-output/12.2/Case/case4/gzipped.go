package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	//定义文件名
	fName := "MyFile.gz"
	//初始化bufio.Reader
	var r *bufio.Reader
	//打开文件
	fi, err := os.Open(fName)
	//判断是否打开成功
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName,
			err)
		os.Exit(1)
	}
	//读取压缩文件
	fz, err := gzip.NewReader(fi)
	//判断是否成功，不成功则直接使用fi打开文件(可能是普通文件)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}
	//循环行，直到文件全被读取
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
