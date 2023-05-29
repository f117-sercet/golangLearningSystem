package main

import "fmt"

func main() {

	// 切片的创建方式

	var slice []int = make([]int, 4)
	fmt.Println(slice) // 默认值为0
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println(slice)

	// 方式二
	var slice2 []string = []string{"1", "2", "1", "4"}
	fmt.Println(slice2)
	fmt.Println(cap(slice2))
	fmt.Println(len(slice2))
}
