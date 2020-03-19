//自己思路的调整版本
package main

import "fmt"

type obj interface{}

func upObj(o obj) obj {
	switch o.(type) {
	case int:
		return o.(int) * 2
	case string:
		return o.(string) + o.(string)
	}
	return o
}

func mapFunc(mf func(obj) obj, list []obj) []obj {
	listlen := len(list)
	for i := 0; i < listlen; i++ {
		list[i] = mf(list[i])
	}
	return list
}

func main() {
	intarray := []obj{1, 2, 3, 4, 5}
	strarray := []obj{"1", "2", "3", "4", "5"}
	fmt.Println(mapFunc(upObj, intarray))
	fmt.Println(mapFunc(upObj, strarray))
}
