package miner

//定义接口，包含三种方法（获取长度，比较，交换）
type Miner interface {
	//获取长度用于循环
	Len() int
	//比较值
	Less(i, temp int) bool
	//交换temp的值
	Swap(i, temp int) int
}

type IntArray []int

func (ia IntArray) Len() int {
	return len(ia)
}

func (ia IntArray) Less(i, temp int) bool {
	return ia[i] < temp
}

func (ia IntArray) Swap(i, temp int) int {
	temp = ia[i]
	return temp
}

func Min(m Miner) int {
	var tmp int
	var LenMin = m.Len()
	for i := 0; i < LenMin; i++ {
		if i == 0 {
			tmp = m.Swap(i, tmp)
		}
		if m.Less(i, tmp) {
			tmp = m.Swap(i, tmp)
		}
	}
	return tmp
}
