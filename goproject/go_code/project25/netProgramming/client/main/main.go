package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8888")

	if err != nil {

		fmt.Println("client dial err= ", err)
	}
	reader := bufio.NewReader(os.Stdin) // 标准输入

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("错误")
	}
	n, err := conn.Write([]byte(line))

	if err != nil {
		fmt.Println("网络出差了")
	}
	fmt.Println("发送", n)

}
