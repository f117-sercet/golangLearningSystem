package main

import "fmt"

func main() {

	//channel 的关闭
	inChan := make(chan int, 10)
	inChan <- 10
	inChan <- 11
	inChan <- 12
	inChan <- 13
	inChan <- 14
	inChan <- 16

	close(inChan) // 关闭Chan

	inChan2 := make(chan int, 100)

	for i := 0; i < 100; i++ {

		inChan2 <- i * 2
	}

	close(inChan2)

	for v := range inChan2 {
		fmt.Println(v)
	}
}
