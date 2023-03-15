package main

import (
	"fmt"
)

func Cal(n1 int, n2 int) int {

	fmt.Println(n1 + n2)
	return n1 + n2
}
func main() {
	n1 := 1
	n2 := 2
	result := Cal(n1, n2)
	fmt.Println("result=", result)

}
