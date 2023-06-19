package main

import "fmt"

// 构造方式
type Person struct {
	Name string
	Age  int64
}

func main() {

	// 方式2：
	p1 := Person{"Marry", 20}

	fmt.Println(p1)
}
