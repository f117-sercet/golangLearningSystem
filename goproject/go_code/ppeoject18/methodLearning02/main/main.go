package main

import "fmt"

// 方法的调用和传参机制原理
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {

	f := 3.14 * c.radius * c.radius
	return f
}
func (c *Circle) area2() float64 {

	c.radius = 10

	result := 3.14 * c.radius * c.radius
	fmt.Printf("返回结果=%v\n", result)
	return result
}
func main() {

	var c Circle
	c.radius = 7
	/*res := c.area()

	fmt.Println(res)
	*/

	area2 := (&c).area2()
	fmt.Printf("area=%v\n", area2)

	var c2 = &Circle{
		1,
	}
	// 在创建结构体指针变量时，把字段名和字段值写在一起，这种写法，不依赖于字段的定义顺序
	fmt.Print(c2)

}
