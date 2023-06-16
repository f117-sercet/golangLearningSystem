package main

import "fmt"

/*
*
如果结构体的字段类型为：指针，slice，和map的零值都是nil，即还没有分配空间
如果需要使用这样的字段,需要先make，才能使用
*/
type Person struct {
	Name   string
	Age    int32
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

func main() {

	var p1 Person
	p1.slice = append(p1.slice, 123)
	p1.slice = append(p1.slice, 456)

	fmt.Println(p1.slice)

}
