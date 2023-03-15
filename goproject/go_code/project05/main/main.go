package main

/**
函数的底层调用机制：
栈区：(基本数据类型，一般说分配到栈区，编译器存在一个逃逸分析)
堆区：引用数据类型一般分配到堆区，编译器存在逃逸分析
代码区：存放代码
*/
import (
	"fmt"
)

func test(n1 int) int {

	n1 = n1 + 1

	fmt.Println(n1)
	return n1

}

func main() {
	n1 := 10
	i := test(n1)
	fmt.Println(i)
}
