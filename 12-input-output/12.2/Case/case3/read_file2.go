package main

import (
	"fmt"
	"os"
)

func main() {
	//通过os.Open打开文件
	file, err := os.Open("D:/Code/GoHomeWork/12-input-output/12.2/Case/case3/products2.txt")
	//判断是否打开有误
	if err != nil {
		panic(err)
	}
	//保证函数执行完后关闭文件
	defer file.Close()
	//定义字符串组用于存储文本内容
	var col1, col2, col3 []string
	//循环读取行
	for {
		var v1, v2, v3 string
		//Fscanln读取以空格分隔的文件（需要提前知道分隔符和列数）
		//如何定义分隔符？
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		// scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
