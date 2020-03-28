package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//定义输入、输出文件
	//inputFile := "D:/Code/GoHomeWork/12-input-output/12.2/Case/case1/input.dat"
	inputFile := "D:/Code/GoHomeWork/12-input-output/12.2/Case/case1/input.dat1"
	outputFile := "D:/Code/GoHomeWork/12-input-output/12.2/Case/case1/input_copy.dat"
	//通过ioutil读取输入文件（ReadFile直接读取所有文件，以[]bytes的形式存储）
	buf, err := ioutil.ReadFile(inputFile)
	//判断是否读取成功，读取失败则退出
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	//将[]bytes转换为string类型
	fmt.Printf("%s\n", string(buf))
	//通过iotuile.WriteFile存储文件
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
