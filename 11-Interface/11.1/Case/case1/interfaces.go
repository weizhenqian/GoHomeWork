//接口的基本使用
package main

import "fmt"

//定义接口
type Shaper interface {
	Area() float32
}

//定义结构体
type Square struct {
	side float32
}

//实现接口中的方法
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	//实现结构体
	sq1 := new(Square)
	//赋值
	sq1.side = 5
	//接口初始化
	var areaIntf Shaper
	//接口赋值
	areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
