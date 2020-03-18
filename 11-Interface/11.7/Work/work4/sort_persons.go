package main

import (
	"fmt"

	"./sort"
)

func main() {
	p1 := sort.Person{Fristname: "Wei", Lastname: "ZhenQian"}
	p2 := sort.Person{Fristname: "Zhang", Lastname: "SiSi"}
	p3 := sort.Person{Fristname: "Liu", Lastname: "JingKai"}
	p4 := sort.Person{Fristname: "Ge", Lastname: "JunHao"}
	persons := sort.Persons{p1, p2, p3, p4}
	/* 	var personer sort.Sorter
	   	personer = persons */
	sort.Sort(persons)
	fmt.Println(persons)
}
