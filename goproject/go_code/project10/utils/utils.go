package utils

import (
	"fmt"
	"strings"
)

/**
函数内部声明/定义的变量叫局部变量，作用域仅限于函数内部
函数挖外部声明/定义的变量叫全局变量，作用域在整个包都有效，如果首字母为大写，则作用域在整个程序都有效
如果一个变量是在一个代码块，比如for/if中，那么这个变量的作用域就在该代码的的代码块

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
	// 返回子串在字符串最后一次出现的index，如没有，返回-1
	index := strings.LastIndex("GO,GOLANG", "GO")
	fmt.Printf("index=%v\n", index)
	return res
}
