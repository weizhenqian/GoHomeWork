/* 类似11.7和示例11.3/4，定义一个包 float64，并在包里定义类型 Float64Array，然后让它实现 Sorter 接口用来对 float64 数组进行排序。
另外提供如下方法：
NewFloat64Array()：创建一个包含25个元素的数组变量（参考10.2）
List()：返回数组格式化后的字符串，并在 String() 方法中调用它，这样就不用显式地调用 List() 来打印数组（参考10.7）
Fill()：创建一个包含10个随机浮点数的数组（参考4.5.2.6）
在主程序中新建一个此类型的变量，然后对它排序并进行测试。 */
package main

import (
	"fmt"

	"./float64"
)

func main() {
	/* var data float64.Float64Array
	data := float64.Float64Array{3.2, 1, 4.4, 6, 3, 9}
	a := float64.Float64Array(data) */
	//var data float64.Float64Array
	data := float64.Float64Array(float64.NewFloat64Array())
	var aer float64.Sorter
	aer = data
	float64.Sort(aer)
	fmt.Println(aer)
}
