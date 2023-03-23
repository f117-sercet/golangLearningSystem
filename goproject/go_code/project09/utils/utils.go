package utils

import "fmt"

var Name1 string
var Name2 string
var Name3 string

/****
defer
在函数中，程序员经常需要创建资源，为了在函数执行完毕后，及时的释放资源，设计者提供延时机制
主要目的：函数执行完毕后，函数释放资源
函数执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈(defer栈)
当函数执行完毕后，再从defer栈中按照陷入后出的方式出栈执行
在defer将语句放入栈中，也会将相关的值拷贝同时入栈

*/

func init() {

	Name1 := 'a'
	Name2 := 'b'
	Name3 := 'C'
	fmt.Println(Name1, Name3, Name2)
}

func Sum(n1 int, n2 int) int {

	defer fmt.Print("ok n1=", n1)
	defer fmt.Println("ok n2", n2)

	res := n1 + n2

	fmt.Println("ok3 res=", res)

	return res
}
