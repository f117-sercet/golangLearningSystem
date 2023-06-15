package main

import (
	"fmt"
)

// 创建结构体
type Cat struct {
	Name  string
	Age   string
	color string
}

func main() {

	fmt.Println("learning")

	var cat Cat
	cat.Name = "1"
	cat.Age = "2"
	cat.color = "red"
	fmt.Println(cat)
}
