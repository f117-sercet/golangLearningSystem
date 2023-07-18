package main

import "fmt"

func main() {

	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Println(intChan)

	i := 10
	intChan <- i
	intChan <- int(20.0)

	i2 := cap(intChan)
	fmt.Println(len(intChan), i2)

	// 取数据
	i3 := <-intChan
	fmt.Println(i3)

}
