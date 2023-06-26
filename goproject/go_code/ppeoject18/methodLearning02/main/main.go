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

}
