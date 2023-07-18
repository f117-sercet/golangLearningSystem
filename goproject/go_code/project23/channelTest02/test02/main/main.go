package main

import (
	"fmt"
)

func putNum(intChan chan int) {

	for i := 0; i < 8000; i++ {
		intChan <- i
	}
	close(intChan)
}
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	// 使用for循环
	//var num int
	var flag bool

	for {
		num, ok := <-intChan
		if !ok { //intChan 取不到
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum 协程因为取不到数据，退出")

	//向退出的管道写入true
	exitChan <- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 4)

	// 开启一个协程， 放入8000个
	go putNum(intChan)

	// 开启协程，判断素数
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {

		// 主线程处理
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
		//取出4个T，关闭
	}()
	for {
		res, ok := <-primeChan

		if ok {
			fmt.Println("素数", res)
		}
	}

	fmt.Println("main线程退出")

}
