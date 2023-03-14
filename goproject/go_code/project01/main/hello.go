package main

import "fmt"

/**
如果给包取了别名，则需要使用别名来访访问该包的函数和变量
在同一包下，不能有相同的函数名和相同的全局变量名，
否则报重复定义
main包只能有一个。编译后生成一个有默认名的可执行文件，在$goPath目录下，可以指定名字和目录
*/

func main() {
	fmt.Println("hello,world")
}
