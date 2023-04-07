package main

import "fmt"

/**
数组和切片
 数组可以多个同一个类型数据，数组也是一种数据结构，在go中，数组是值类型 增加代码可维护性

*/

func main() {

	// 定义数组
	var hes [6]float64 // 数组名的地址就是数组的首地址
	var sum float64 = 0.0
	hes[0] = 3.0
	hes[1] = 4.0
	hes[2] = 5.0
	hes[3] = 6.0
	hes[4] = 7.0
	hes[5] = 8.0

	for i := 0; i < len(hes); i++ {

		sum += hes[i]
		fmt.Println(sum)

	}
	fmt.Println("hes的地址=p\n", &hes[0])
	fmt.Println("hes[0]的地址=v\n", &hes[0])

}
