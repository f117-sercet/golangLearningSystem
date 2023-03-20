package main

import (
	"fmt"
	//"go_code/project07/utils"
)

/**
每一个源文件都可以包含一个init函数，改函数会在main函数前执行，被go运行框架调用，也就是说，init会在main函数前调用;
1.如果一个文件同时包含全局变量定义，ini t函数和main函数，则执行的流程是变量定义->init函数->main函数
2.init 函数，通常可以在init函数中完成初始化工作
*/

var age = test()

func test() int {
	fmt.Println("test")
	return 90
}
func init() {

	fmt.Print("init()\t")
}

func main() {

	fmt.Println("a", age)
	//fmt.Println("Age=", utils.Age, "Name=", utils.Name)

	/**
	执行流程
	1.utils.go的变量定义（前提utils没有引入其他包）
	2.utils.go的init函数
	3.main.go的变量定义
	4.main.go的init函数
	5.main.go的main函数
	*/

}
