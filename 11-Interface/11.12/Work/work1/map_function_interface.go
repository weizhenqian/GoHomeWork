/* 在练习 7.13 中我们定义了一个 map 函数来使用 int 切片 （map_function.go）。

通过空接口和类型断言，现在我们可以写一个可以应用于许多类型的 泛型 的 map 函数，为 int 和 string 构建一个把 int 值加倍和连接字符串值的 map 函数 mapFunc。

提示：为了可读性可以定义一个 interface{} 的别名，比如：type obj interface{} */
package main

import "fmt"

type obj interface{}

func main() {
	//此处在自己编写的时候，出现了问题
	//自己是在main函数外定义了两个函数，但是出现的问题，就是值传递的问题。
	mf := func(o obj) obj {
		switch o.(type) {
		case int:
			return o.(int) * 2
		case string:
			return o.(string) + o.(string)
		}
		return o
	}

	intarray := []obj{1, 2, 3, 4, 5}
	res1 := mapFunc(mf, intarray)
	fmt.Println(res1)
	stringarray := []obj{"1", "2", "3", "4", "5"}
	res2 := mapFunc(mf, stringarray)
	fmt.Println(res2)
}

//对比map_function，此处通过[]obj（空接口切片）替代了int/string具体的类型
//函数的接收值和返回值也均为空接口
func mapFunc(mf func(obj) obj, ia []obj) []obj {
	lia := len(ia)
	for i := 0; i < lia; i++ {
		ia[i] = mf(ia[i])
	}
	return ia
}
