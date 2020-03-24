//稍微改变练习 11.9，允许 mapFunc 接收不定数量的 items。
//我理解不定数量的items，是1-n个[]obj，所以采用的思路是将[]obj新增到一个[]obj里面，然后调用函数
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

func mapFunc(mf func(obj) obj, list ...[]obj) []obj {
	//获取到list的长度，判断包含多少个[]obj
	listlen := len(list)
	var listobjs []obj
	for i := 0; i < listlen; i++ {
		//将切片通过append函数进行追加，此处注意，如果append第二个参数为切片的话，需要在切片名称的后面加上...
		listobjs = append(listobjs, list[i]...)
	}
	loslen := len(listobjs)
	for i := 0; i < loslen; i++ {
		listobjs[i] = mf(listobjs[i])
	}
	return listobjs
}

func main() {
	intarray := []obj{1, 2, 3, 4, 5}
	intarray2 := []obj{6, 7, 8, 9, 10}
	intarray3 := []obj{11, 12, 13, 14, 15}
	//strarray := []obj{"1", "2", "3", "4", "5"}
	fmt.Println(mapFunc(upObj, intarray, intarray2, intarray3))
	//fmt.Println(mapFunc(upObj, strarray))
}
