package main

import "fmt"

/**
每一个源文件都可以包含一个init函数，改函数会在main函数前执行，被go运行框架调用，也就是说，init会在main函数前调用;
1.如果一个文件同时包含全局变量定义，init函数和main函数，则执行的流程是变量定义->init函数->main函数
*/

func init() {

	fmt.Print("init()")
}

func main() {

	fmt.Println("main()")

}
