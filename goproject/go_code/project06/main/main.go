package main

import "fmt"

/**
1)在go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变凉了，通过该变量可以对函数进行调用
2) 函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用
3)为了简化数据类型定义，go支持自定义数据类型 基本语法：type 自定义数据类型名，数据类型 // ：相当于一个别名
4)支持对函数返回值命名
5)使用_标识符，忽略返回值
6)支持可变参数
*/

func getSum(n1 int, n2 int) int {

	return n1 + n2

}

// 函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}
func Cal1(n1 int, n2 int) (sum int, sub int) {

	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {

	a := getSum
	fmt.Print("a的类型%T，getSum类型是%T\n", a, getSum)

	res := a(10, 40)

	fmt.Println("res", res)

	res2 := myFun(getSum, 50, 60)

	fmt.Println("res2=", res2)

	type MyInt int // 给int取别名,但是是两个不同的类型
	var n1 MyInt = 1
	var n2 int
	n2 = int(n1)
	fmt.Println(n1, n2)

	// 参数类型
	sum, sub := Cal1(10, 20)

	fmt.Println(sum, sub)

}
