/*a). 继续扩展程序，定义类型 Triangle，让它实现 AreaInterface 接口。通过计算一个特定三角形的面积来进行测试（三角形面积=0.5 (底 高)）

b). 定义一个新接口 PeriInterface，它有一个 Perimeter 方法。让 Square 实现这个接口，并通过一个 Square 示例来测试它。*/
package main

import "fmt"

type AreaInterface interface {
	Area() float32
}

type Triangle struct {
	hight    float32
	baseline float32
}

func (t Triangle) Area() float32 {
	return 0.5 * t.hight * t.baseline
}

type Square struct {
	side float32
}

type PeriInterface interface {
	Perimeter() float32
}

func (s Square) Perimeter() float32 {
	return 4 * s.side
}

func main() {
	//a
	triangle := Triangle{3, 5}
	var areaInterface AreaInterface
	areaInterface = triangle
	fmt.Println(areaInterface.Area())
	//b
	square := Square{4}
	var perinterface PeriInterface
	perinterface = square
	fmt.Println(perinterface.Perimeter())
}
