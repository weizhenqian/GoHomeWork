//类型 Rectangle 也实现了 Shaper 接口。接着创建一个 Shaper 类型的数组，迭代它的每一个元素并在上面调用 Area() 方法，以此来展示多态行为
//此案例，记录了两个结构体，通过定义接口类型的数组，迭代调用的过程
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
func main() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
