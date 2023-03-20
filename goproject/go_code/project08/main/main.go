package main

import "fmt"

/**
匿名函数
Go支持匿名函数，如果我们某个函数只是希望使用一次，剋以考虑使用匿名函数，匿名函数也可以实现多次调用
匿名函数使用方式
1，在定义匿名函数时就直接调用，只能调用一次
2，将匿名函数赋给一个变量(函数变量)，再通过该变量来调用匿名函数
全局匿名函数
如果将匿名函数赋给一个全局变量，那么这个匿名函数就成为一个全局匿名函数，可以在程序有效
*/

/*
*
全局匿名函数
*/
var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 + 2*n2
	}
)

func main() {

	// 定义时直接调用

	res1 := func(n1 int, n2 int) int {

		return n1 + n2
	}(10, 20)

	fmt.Println("res1的值\n", res1)

	// 将匿名函数func(n1 int,n2 int ) int 赋给变量a
	// 则 a的数据结构类型就是函数类型，此时，我们通过a完成调用
	// 此时的a可以反复调用
	a := func(n1 int, n2 int) int {

		return n1 - n2
	}
	res2 := a(10, 20)

	fmt.Println("res2的值\n", res2)

	// 全局匿名函数调用
	res3 := Fun1(10, 20)
	fmt.Println("res3的值\n", res3)
}
