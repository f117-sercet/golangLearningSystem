package main

import (
	"fmt"
	"runtime"
)

func main() {

	// cpu个数
	numCPU := runtime.NumCPU()

	fmt.Println("Cpu:", numCPU)

	//设置使用cpu的个数
	gomaxprocs := runtime.GOMAXPROCS(numCPU - 1)
	fmt.Println("goMaxProcess", gomaxprocs)
	//lock := sync.Mutex{}
	//lock.Lock()
}
