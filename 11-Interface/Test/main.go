//一直搞不明白*和&的引用关系，此处编写test用于测试
package main

import (
	"fmt"
	"reflect"
)

func main() {
	//定义了p，p的类型是指向int类型的指针
	var p *int
	fmt.Println(p)
	fmt.Println(reflect.TypeOf(p))
	//定义了i为int类型，&i取i的指针地址，赋值给p
	i := 42
	p = &i
	fmt.Println(p)
	fmt.Println(reflect.TypeOf(p))
	//通过*p可以获取到指针底层值
	fmt.Println(*p)
	//对指针底层值进行修改，i的值也会改变
	*p = 21
	fmt.Println(i)
}

//总结：*如果在类型上，则是指向类型的指针；如果在变量上，则是获取底层值。&一般作用于变量，用于获取指针地址。
