/*似11.7和示例11.3/4，定义一个包 float64，并在包里定义类型 Float64Array，然后让它实现 Sorter 接口用来对 float64 数组进行排序。

另外提供如下方法：

NewFloat64Array()：创建一个包含25个元素的数组变量（参考10.2）
List()：返回数组格式化后的字符串，并在 String() 方法中调用它，这样就不用显式地调用 List() 来打印数组（参考10.7）
Fill()：创建一个包含10个随机浮点数的数组（参考4.5.2.6）
在主程序中新建一个此类型的变量，然后对它排序并进行测试。*/
package float64

import (
	"math/rand"
)

type Float64Array []float64

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (f Float64Array) Len() int           { return len(f) }
func (f Float64Array) Less(i, j int) bool { return f[i] > f[j] }
func (f Float64Array) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func NewFloat64Array() []float64 {
	a := []float64{}
	for i := 0; i < 25; i++ {
		a = append(a, rand.Float64())
	}
	return a
}
