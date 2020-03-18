/* 继续 练习11.2，它中添加一在个 gI 函数，它不再接受 Simpler 类型的参数，而是接受一个空接口参数。然后通过类型断言判断参数是否是 Simpler 类型。最后在 main 使用 gI 取代 fI 函数并调用它。确保你的代码足够安全。 */
package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Any interface {
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

func gI(any Any) {
	switch v := any.(type) {
	case *Square:
		fmt.Printf("any %v is a Square type", v)
	case Rectangle:
		fmt.Printf("any %v is a Rectangle type", v)
	case Circle:
		fmt.Printf("any %v is a Circle type", v)
	}
}

func main() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	c := Circle{2.5}
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	/* 	shapes := []Shaper{r, q, c}
	   	fmt.Println("Looping through shapes for area ...")
	   	for n, _ := range shapes {
	   		fmt.Println("Shape details: ", shapes[n])
	   		fmt.Println("Area of this shape is: ", shapes[n].Area())
	   	} */
	var val Any
	val = r
	gI(val)
	val = q
	gI(val)
	val = c
	gI(val)
}
