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

	// 遍历
	// for循环遍历
	for i := 0; i < len(slice); i++ {

		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}

	//for-range 遍历切片
	for i, v := range slice2 {

		fmt.Printf("i=%v v=%v \n", i, v)
	}
	// slice 切片学习
}
