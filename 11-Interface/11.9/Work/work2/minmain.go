package main

import (
	"fmt"

	"./miner"
)

func main() {
	intarray := miner.IntArray{1, 3, 5, 6, 8, 4}
	var m miner.Miner
	m = intarray
	fmt.Println(miner.Min(m))
}
