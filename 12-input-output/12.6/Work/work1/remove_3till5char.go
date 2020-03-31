/* 下面的代码有一个输入文件 goprogram，然后以每一行为单位读取，从读取的当前行中截取第 3 到第 5 的字节写入另一个文件。然而当你运行这个程序，输出的文件却是个空文件。找出程序逻辑中的 bug，修正它并测试。 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("goprogram")
	outputFile, _ := os.OpenFile("goprogramT", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		fmt.Println(inputString)
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}
		outputString := string(inputString[:]) + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
		//因为使用了缓存，需要Flush
		outputWriter.Flush()
	}
	fmt.Println("Conversion done")
}
