/* flag.Parse() 扫描参数列表（或者常量列表）并设置 flag, flag.Arg(i) 表示第i个参数。Parse() 之后 flag.Arg(i) 全部可用，flag.Arg(0) 就是第一个真实的 flag，而不是像 os.Args(0) 放置程序的名字。

flag.Narg() 返回参数的数量。解析后 flag 或常量就可用了。
flag.Bool() 定义了一个默认值是 false 的 flag：当在命令行出现了第一个参数（这里是 “n”），flag 被设置成 true（NewLine 是 *bool 类型）。flag 被解引用到 *NewLine，所以当值是 true 时将添加一个 newline（”\n”）。

flag.PrintDefaults() 打印 flag 的使用帮助信息，本例中打印的是：-n=false: print newline */
package main

import (
	"flag" // command line option parser
	"os"
)

var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool
const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
