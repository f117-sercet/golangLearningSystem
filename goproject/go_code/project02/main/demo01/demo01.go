package main

import "fmt"

func getVal(num1 int, num2 int) (int, int) {

	sum := num1 + num2
	sub := num2 - num1

	return sum, sub
}

func main() {

	sub, sum := getVal(30, 30)
	fmt.Println("sum=", sum, "sub=", sub)
	sum2, nice := getVal(10, 30) // 只取出第一个返回值
	fmt.Println("sum=", sum2, "nice=", nice)
}
