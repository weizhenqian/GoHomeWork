//用示例 13.3 中的编码模式通过整数除以 0 触发一个运行时 panic。
package main

import (
	"fmt"
)

var err interface{}

func except(a, b int) int {
	return a / b
}
func test(a, b int) (result int, err interface{}) {
	defer func() {
		if e := recover(); e != nil {
			err = e
		}
	}()
	result = except(a, b)
	return
}
func main() {
	fmt.Printf("Calling test\r\n")
	res, e := test(2, 0)
	if e == nil {
		fmt.Printf("res is %v \r\n", res)
	} else {
		fmt.Printf("res is %v \r\n erros is %v \r\n", res, e)
	}

	fmt.Printf("Test completed\r\n")
}
