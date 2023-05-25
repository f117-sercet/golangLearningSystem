package main

import "fmt"

// 错误处理机制：
// 1.在默认情况下，当发生错误后， 程序就会退出
// 2.如果我们希望，当发生错误后，可以捕获到错误，并进行处理，保证程序可以继续执行，还可以捕获到错误后，给管理员一个提示
/******************************错误处理基本说明**************************************************************/
/**
1.go语言追求简洁优雅，所以，Go语言不支持传统的try....catch finally 这种处理
2.Go语言引入处理的方式为：defer1，panic，recover
3.这个异常的使用场景可描述为：Go可以跑出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理
4.可以将错误提前预警，让代码更加健壮
*/
/*****************************自定义错误***************************************************************************/
/**
Go程序中，也支持自定义错误，使用errors.New和panic内置函数
1.error.new("错误说明")，会返回一个error类型的值，表示一个错误
2.panic内置函数，接受一个interface{}类型的值作为参数，可以接收error类型的变量，输出错误信息，并退出程序
*/
func test() {
	// 使用defer+ recover来捕获和处理异常
	defer func() {
		err := recover() // recover()为内置函数，可以捕获异常
		if err != nil {  // 捕获异常
			fmt.Println("err=", err)

		}
		num1 := 10
		num2 := 0
		res := num1 / num2
		fmt.Print(res)
	}()
}
func main() {

	test()
}