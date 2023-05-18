package main

func main() {

	// 长度固定，不能动态变化
	var arr01 [3]int

	arr01[0] = 1
	arr01[1] = 30
	// 报错
	//arr01[2] = 1.1
	// 数组越界
	//arr01[3] = 4
}
