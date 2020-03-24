package stack

import (
	"errors"
)

//限定栈结构的长度
const LIMIT = 10

//定义Stack结构体
type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

//先进先出或者先进后出
func (s *Stack) Pop() (interface{}, error) {
	stack := *s
	//首先判断Stack的长度，如果为空则返回错误
	if len(stack) == 0 {
		return nil, errors.New("stack is empty")
	}
	//获取到Stack的具体值
	//先进先出的栈
	//将切片进行缩减，可以通过s=s[:len(s)-1]的方法缩减
	/* end := stack[len(stack)-1]
	*s = stack[:len(stack)-1]
	return end, nil */
	//先进后出的栈
	//将切片进行缩减，可以通过s=s[1:]的方法缩减
	head := stack[0]
	*s = stack[1:]
	return head, nil
}
