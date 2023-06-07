package main

import "fmt"

func main() {

	var slice []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	slice = arr[:]
	var slice2 = slice
	slice2[0] = 10

	fmt.Println("slice2", slice2)
	fmt.Println("slice", slice)
	fmt.Println("arr", arr)

	// []byte 以字节来处理，而一个汉字，是三个字节，因此会出现乱码
	str := "hello@atguigu"
	slice3 := str[6:]
	fmt.Println("slice3", slice3)
}
