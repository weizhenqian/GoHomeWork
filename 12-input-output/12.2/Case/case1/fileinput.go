package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//此处通过绝对路径进行的文件打开，不知道如何通过相对路径打开
	inputFile, inputError := os.Open("D:/Code/GoHomeWork/12-input-output/12.2/Case/case1/input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	//函数结束时，关闭打开着的文件
	defer inputFile.Close()
	//初始化bufio提供的读取器
	inputReader := bufio.NewReader(inputFile)
	for {
		//当文件读取结束时，第二个返回值为EOF，将其赋值于readerError
		//ReadString、ReadBytes、ReadLine三种方法都可以获取
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}
