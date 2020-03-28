package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var outputWriter *bufio.Writer
	// var outputFile *os.File
	// var outputError os.Error
	// var outputString string
	//以读写方式打开文件，不存在则创建，文件权限为0666
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	//打开失败则报错
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	//关闭文件
	defer outputFile.Close()
	//定义缓存，定义写入内容，循环写入
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	//将缓存内容写入文件
	outputWriter.Flush()
}
