/*继续 10.3 中的练习 point_methods.go，定义接口 Magnitude，它有一个方法 Abs()。让 Point、Point3 及Polar 实现此接口。通过接口类型变量使用方法做point.go中同样的事情。*/
package main

import (
	"fmt"
	"math"
)

type Magnitude interface {
	Abs() float64
}

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
	/*
		point := Point{3, 4}
		polar := Polar{3, 4, 5}
		fmt.Println(point.Abs())
		fmt.Println(polar.Abs())
		fmt.Println(point.Scale(3))
		fmt.Println(polar.Scale(4))
	*/
	point := Point{3, 4}
	polar := Polar{3, 4, 5}
	var magnitude Magnitude
	magnitude = point
	fmt.Println(magnitude.Abs())
	magnitude = polar
	fmt.Println(magnitude.Abs())
}
