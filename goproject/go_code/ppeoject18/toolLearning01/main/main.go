package main

import "fmt"

type A struct {
	Num int
}

func (a A) test() {
	fmt.Println("test", a.Num)
}

func main() {

	var a A
	a.Num = 1
	a.test()
}
