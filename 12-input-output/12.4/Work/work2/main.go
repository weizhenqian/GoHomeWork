/* 测试自己是否掌握了flag函数包，编写一个案例。
简易版open-falcon-agent */
package main

import (
	"flag"
	"os"
)

var (
	check   *bool
	config  *string
	version *bool
)

func init() {
	check = flag.Bool("check", false, "check agent health")
	config = flag.String("f", "", "agent config file")
	version = flag.Bool("v", false, "print agent version")
}

func main() {
	flag.Parse()
	if *check {
		os.Stdout.WriteString("agent is alived")
		os.Exit(1)
	}

	if *version {
		os.Stdout.WriteString("version:1.0.0")
		os.Exit(1)
	}

	if *config != "" {
		*config = "agent config is:" + *config
		os.Stdout.WriteString(*config)
		os.Exit(1)
	}
}
