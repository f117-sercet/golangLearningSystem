package main

import "fmt"

/**
闭包：
闭包就是一个函数和与其相关的引用环境组合的一个整体(实体)
*/

/*
*
1.闭包的说明:返回的是一个匿名函数，但是这个匿名函数引用到函数外的n,形成引用关系，构成闭包
2.闭包是类，函数是操作，n是字段
3.当反复调用f时，n是初始化一次，每调用一次，进行累加，
*/
func AddUpper() func(int) int {

	var n int = 10
	var str = "hello"
	return func(x int) int {
		str += string(36)
		fmt.Println("str\n" + str)
		n = n + x
		return n
	}

}

func main() {

	// 累加器
	f := AddUpper()
	fmt.Print(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
