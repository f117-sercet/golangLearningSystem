package main

import "fmt"

func main() {

	// 二维数组的演示案例
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], "")
		}
		fmt.Println()
	}
	fmt.Println()

	// 声明方式2
	var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr2=", arr2)

	// 声明方式3
	var arr3 [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3=", arr3)

	// 声明方式4
	var arr4 = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr4=", arr4)

	// 声明方式5
	var arr5 = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr5=", arr5)

}
