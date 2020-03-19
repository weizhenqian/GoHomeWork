//同map_function_interface，此处尝试自己的思路解决问题
//这一版，是通过mapFunc的时候进行了类型判断，虽然也实现了功能，但是有些偏差
//优化思路：可以将upInt和upString合并为upObj，在upObj中进行类型判断，通过mapFunc进行调用
package main

import "fmt"

type obj interface{}

func upInt(o obj) obj {
	return o.(int) * 2
}

func upString(o obj) obj {
	return o.(string) + o.(string)
}

func mapFunc(list []obj) []obj {
	listlen := len(list)
	for i := 0; i < listlen; i++ {
		switch list[i].(type) {
		case int:
			list[i] = upInt(list[i])
		case string:
			list[i] = upString(list[i])
		}
	}
	return list
}

func main() {
	intarray := []obj{1, 2, 3, 4, 5}
	strarray := []obj{"1", "2", "3", "4", "5"}
	fmt.Println(mapFunc(intarray))
	fmt.Println(mapFunc(strarray))
}
