package main

import (
	"fmt"
	"os"
)

// 文件
func main() {

	// filie叫法
	/**
	1.叫对象
	2.叫file 指针
	3，叫file句柄
	*/
	// 打开文件
	file, err := os.Open("C:/Users/LENOVO/Desktop/完成度.xls")
	name := file.Name()
	fmt.Println(*file, err, name)

	// 关闭文件
	file.Close()

}
