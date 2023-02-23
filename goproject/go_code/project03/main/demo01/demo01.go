package main

import "fmt"

func main() {

	var n1 int = 10
	var n2 int = 40
	var max int
	if n1 > n2 {

		max = n1

	} else {

		max = n2
	}
	fmt.Println("max=", max)

	// 输入变量
	var name string
	var age byte
	var sal float32
	var isPass bool

	fmt.Println("请输入姓名")

	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)

	fmt.Println("请输入是否")
	fmt.Scanln(&isPass)

	fmt.Println("名字 \n age %v \n sal %v \n isPass %v \n", name, age, sal, isPass)

}
