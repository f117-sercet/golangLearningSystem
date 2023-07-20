package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {

	defer conn.Close()
	for {
		// 创建新的切片
		slice := make([]byte, 1024)

		// 服务器在等待客户端的输入
		n, err := conn.Read(slice)

		if err != nil {
			fmt.Println("服务器端出错", err)
			return
		}
		fmt.Printf(string(slice[:n]))

	}
}

func main() {

	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp", "localhost:8888")

	if err != nil {
		fmt.Println("error=", err)
		return
	}

	defer listen.Close() // 延时关闭

	for {
		// 等待连接
		fmt.Println("等待客户端连接")
		conn, _ := listen.Accept()

		fmt.Println("conn", conn, "客户端信息", conn.RemoteAddr().String())

		go process(conn)
	}

	//fmt.Println("建立v成功", listen)

}
