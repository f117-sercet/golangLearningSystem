package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// filie叫法
	/**
	1.叫对象
	2.叫file 指针
	3，叫file句柄
	*/
	// 打开文件
	file, err := os.Open("C:/Users/LENOVO/Desktop/demo.txt")
	name := file.Name()
	fmt.Println(*file, err, name)

	// 关闭文件
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
	fmt.Println("文件读取结束")

	// 读取文件的第二种方式
	readFile, err := os.ReadFile("C:/Users/LENOVO/Desktop/demo.txt")
	fmt.Println(string(readFile))

	// 创建文件
	create, err := os.Create("C:/Users/LENOVO/Desktop/demoTest01.txt")
	fmt.Println(create)

	defer file.Close()
	// 写文件
	//err := os.WriteFile("C:/Users/LENOVO/Desktop/demoTest01.txt", nil, os.ModeAppend)

	//newReader := bufio.NewReader(file)
	//newReader.Flush()

	// 名命令行参数
}
