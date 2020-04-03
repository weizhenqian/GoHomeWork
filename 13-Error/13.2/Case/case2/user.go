package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func check() {
	fmt.Printf("user is %v", user)
	if user == "" {
		panic("Unknown user: no value for $USER")
	}
}

func main() {
	check()
}
