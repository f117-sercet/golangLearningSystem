package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) Sayok() {
	fmt.Println("A Sayok", a.Name)
}
func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
}

// 继承
func main() {

	var b B
	b.A.Name = "tom"
	b.A.age = 18
	b.A.Sayok()
	b.hello()
}
