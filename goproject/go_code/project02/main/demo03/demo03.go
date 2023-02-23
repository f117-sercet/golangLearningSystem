package main

import "fmt"

/*
**
字符介绍
*/
func main() {

	var c1 byte = 'a'
	var c2 byte = '0'

	fmt.Println("c1=", c1)
	fmt.Println("c2", c2)

	var c4 int = 2269
	fmt.Println("c4=%c\n", c4)

	var a int

	var b float32

	var c float64

	var isMarried bool
	isMarried = true
	var name string

	//这里的%v 表示按照变量的值输出
	fmt.Println("a=%d", a)
	fmt.Println("b=%d", b)
	fmt.Println("c=%v", c)
	fmt.Println("isMarried=%d", isMarried)

	fmt.Println("name=", name)

}
