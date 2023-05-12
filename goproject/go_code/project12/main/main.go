package main

import "fmt"

/**
数组和切片
 数组可以多个同一个类型数据，数组也是一种数据结构，在go中，数组是值类型 增加代码可维护性

1)数组的地址可以通过数组名来获取&intArr

2)数组的第一个元素的地址，就是数组的首地址

3)元素连续分布，各个元素的地址大小间隔取决于数组的类型决定的

***数组的初始化
方式一：常规遍历

    for i:=0 i<len(nums);i++{}

方式二：for-range结构遍历

基本语法：

    for index ,value:=range array01{

        .........
    }


说明：

1）第一个返回值index是数组的下标

2）第二个value是在该下标位置的值

3）他们都是仅在for循环内部可见的局部变量

4）遍历数组元素的时候，如果不想使用下标index，可以直接把下标index标为下划线__

5）index和value的名称不i是固定的，即程序员可以自行指定，一般命名为index和value




*/

func main() {

	// 定义数组
	var hes [6]float64 // 数组名的地址就是数组的首地址
	var sum float64 = 0.0
	hes[0] = 3.0
	hes[1] = 4.0
	hes[2] = 5.0
	hes[3] = 6.0
	hes[4] = 7.0
	hes[5] = 8.0

	for i := 0; i < len(hes); i++ {

		sum += hes[i]
		fmt.Println(sum)

	}
	fmt.Println("hes的地址=p\n", &hes[0])
	fmt.Println("hes[0]的地址=v\n", &hes[0])

	// 另外的遍历方法
	heros := [...]string{"宋", "吴", "璐"}
	for i, v := range heros {
		fmt.Println("i=%v v=%v\n", i, v)
	}

}
