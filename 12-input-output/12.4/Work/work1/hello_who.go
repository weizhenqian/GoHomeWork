// os_args.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var who string
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}
