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
}

// 继承
func main() {

}
