/*定义一个接口 Simpler，它有一个 Get() 方法和一个 Set()，Get()返回一个整型值，Set() 有一个整型参数。创建一个结构体类型 Simple 实现这个接口。

接着定一个函数，它有一个 Simpler 类型的参数，调用参数的 Get() 和 Set() 方法。在 main 函数里调用这个函数，看看它是否可以正确运行。*/
package main

import "fmt"

type Simper interface {
	Get() int
	Set(int)
}

type Simpe struct {
	id int
}

func (s *Simpe) Get() int {
	return s.id
}

//注意此处，如果为值引用，则无法改变值。需要设置为指针引用。
func (s *Simpe) Set(a int) {
	s.id = a
}

func showSimper(s Simper) {
	fmt.Print(s.Get())
}

func changeSimper(s Simper, a int) {
	s.Set(a)
}

func main() {
	//此处为初始化一个名为simpe的指针，指向Simpe结构体
	simpe := &Simpe{1}
	var simper Simper
	simper = simpe
	/*fmt.Print(simper.Get())
	simper.Set(2)
	fmt.Print(simper.Get())*/
	showSimper(simper)
	changeSimper(simper, 2)
	showSimper(simper)
}
