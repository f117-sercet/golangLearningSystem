package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("命令行参数", len(os.Args))

	for i, v := range os.Args {

		fmt.Printf("args[%v]=%v\n", i, v)

	}

	// flag 命令解析

	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "p", "", "密码")
	flag.StringVar(&host, "host", "", "主机名")
	flag.IntVar(&port, "port", 3306, "端口")
	flag.Parsed()
	fmt.Println(user, pwd, host, port, host)
}
