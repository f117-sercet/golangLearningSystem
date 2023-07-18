package main

import (
	"fmt"
)

// writeData
func writeData(intChan chan int) {

	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Printf("writeData 读到数据=%v\n", i)
	}
	close(intChan)
}

// readData
func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}

		fmt.Printf("readData 读到数据=%v\n", v)
	}
	exitChan <- true
	close(exitChan)
}

func main() {

	// 创建两个管道
	intChan := make(chan int, 100)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {

		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
