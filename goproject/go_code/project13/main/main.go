package main

import (
	"fmt"
)

func main() {

	var intArr [5]int = [...]int{1, 2, 3, 4, 5}

	slice := intArr[1:3]
	// 声明一个切片
	// slice:=intArrp[1:3]
	///1.slice 就是切片名
	// intArr【1:3】 表示slice引用到intArr这个数组(起始下标《变量名<终止下标)

	fmt.Println("intArr=", intArr)
	fmt.Println("slice的元素=", slice)
	fmt.Println("slice的元素个数 = ", slice)
	fmt.Println("slice的元素个数 = ", len(slice))
	fmt.Println("slice 的容量=", cap(slice))

	// 切片内存分析
}
