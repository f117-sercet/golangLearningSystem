package main

import "fmt"

// 接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量
type AInterface interface {
	Say()
}
type Stu struct {
	Name string
}
type integer int

func (i integer) Say() {
	fmt.Println("integer Say i=", i)
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

func main() {

	var stu Stu
	var a AInterface = stu

	a.Say()

	var i integer = 10
	var b AInterface = i
	//只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型。
	b.Say()
}
