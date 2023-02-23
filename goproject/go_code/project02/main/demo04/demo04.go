package main

import (
	"fmt"
	// "goproject/src/go_code/project02/main/demo04/model"
)

// 演示golan指针类型
func main() {

	var i int = 10

	// i的地址是什么
	fmt.Println("i的地址=", &i)

	var ptr *int = &i

	fmt.Println("ptr=", ptr)
	fmt.Println("ptr的地址=", &ptr)
	fmt.Println("ptr的指向地址=", *ptr)
	//fmt.Println()

}
