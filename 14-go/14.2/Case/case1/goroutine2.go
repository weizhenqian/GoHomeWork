package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	//此处sleep，为了保证程序运行完整
	time.Sleep(1e9)
}
func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}
func getData(ch chan string) {
	var input string
	for {
		if input = <-ch; input == "London" {
			fmt.Printf("%s is dead \r\n", input)
		} else {
			fmt.Printf("%s is ok \r\n", input)
		}
	}
}
