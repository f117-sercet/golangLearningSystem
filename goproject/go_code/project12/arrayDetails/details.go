package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 长度固定，不能动态变化
	var arr01 [3]int

	arr01[0] = 1
	arr01[1] = 30
	// 报错
	//arr01[2] = 1.1
	// 数组越界
	//arr01[3] = 4

	//数组使用步骤
	/**
	1.声明并且开辟空间
	2.赋值
	3.使用数组
	4.数组的下标是从0开始的
	5.Go的数组属值类型，在默认的情况下是值传递，因此会值拷贝
	6.若想在其他函数中，去修改原来的数组，可以使用引用传递（指针方式）
	数组
	*/
	var intArr5 [5]int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr5); i++ {
		intArr5[i] = rand.Intn(100)
	}
	fmt.Println(intArr5)
}
