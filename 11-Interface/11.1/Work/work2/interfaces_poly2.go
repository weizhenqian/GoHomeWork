/*a) 扩展 interfaces_poly.go 中的例子，添加一个 Circle 类型*/
/*b) 使用一个抽象类型 Shape（没有字段） 实现同样的功能，它实现接口 Shaper，然后在其他类型里内嵌此类型。扩展 10.6.5 中的例子来说明覆写。*/
package main

import "fmt"

type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return 2 * 3.14 * c.radius
}

func main() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	c := Circle{2.5}
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q, c}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
