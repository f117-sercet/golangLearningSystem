package main

import (
	"fmt"
)

func main() {

	intChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)

	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	for {

		select {
		//若intchan一直没有关闭，不会一直阻塞，会自动到下一个case匹配
		case v := <-intChan:
			fmt.Println("读数", v)
			return
		case v := <-stringChan:
			fmt.Println("读数", v)
			return
		default:
			fmt.Println("完犊子")
			return
		}

	}

}
