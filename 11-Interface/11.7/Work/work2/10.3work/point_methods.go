/*使用坐标 X、Y 定义一个二维 Point 结构体。同样地，对一个三维点使用它的极坐标定义一个 Polar 结构体。实现一个 Abs() 方法来计算一个 Point 表示的向量的长度，实现一个 Scale 方法，它将点的坐标乘以一个尺度因子（提示：使用 math 包里的 Sqrt 函数）（function Scale that multiplies the coordinates of a point with a scale factor）。*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Polar struct {
	X float64
	Y float64
	Z float64
}

func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p Polar) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

func (p Point) Scale(s float64) (float64, float64) {
	X := p.X * s
	Y := p.Y * s
	return X, Y
}

func (p Polar) Scale(s float64) (float64, float64, float64) {
	X := p.X * s
	Y := p.Y * s
	Z := p.Z * s
	return X, Y, Z
}

func main() {
	point := Point{3, 4}
	polar := Polar{3, 4, 5}
	fmt.Println(point.Abs())
	fmt.Println(polar.Abs())
	fmt.Println(point.Scale(3))
	fmt.Println(polar.Scale(4))
}
