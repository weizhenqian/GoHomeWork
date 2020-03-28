package main

import (
	"os"
	"strconv"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

//代码实现[]byte转换为string
func convert(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}

func (p *Page) save() {
	//定义文件名
	outputFile := p.Title
	//打开文件，不存在则创建
	f, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0)
	//如果打开失败则报错
	if err != nil {
		panic(err)
	}
	//定义写入内容
	outputContent := p.Body
	//注意：此处Body的类型为[]byte，如何将byte以string的形式写入到文件？
	//尝试使用string转换，发现由于编码原因123456不可见。
	//fmt.Println(string(outputContent[:]))
	f.WriteString(convert(outputContent[:]))
}

/* func load(p *Page) {

} */

func main() {
	content := []byte{1, 2, 3, 4, 5, 6}
	page := Page{"index.html", content}
	page.save()
}
