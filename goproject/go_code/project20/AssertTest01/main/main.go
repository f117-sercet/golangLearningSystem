package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {

	// 类型断言
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	b = a.(Point)
	fmt.Println(b)

}
