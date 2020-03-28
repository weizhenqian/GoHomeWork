package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	productsfile, err := os.Open("D:/Code/GoHomeWork/12-input-output/12.2/Work/work1/products.txt")
	if err != nil {
		panic(err)
	}
	defer productsfile.Close()
	productsread := bufio.NewReader(productsfile)
	var title, price, quantity []string
	for {
		productstring, productserr := productsread.ReadString('\n')
		//对读取到的字符串进行切割
		s := strings.Split(productstring, ";")
		title = append(title, s[0])
		price = append(price, s[1])
		quantity = append(quantity, s[2])
		if productserr == io.EOF {
			break
		}
	}
	fmt.Println(title)
	fmt.Println(price)
	fmt.Println(quantity)
}
