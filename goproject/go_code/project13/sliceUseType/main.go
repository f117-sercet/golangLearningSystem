package main

import "fmt"

func main() {

	// 切片的创建方式

	var slice []int = make([]int, 4)
	fmt.Println(slice) // 默认值为0
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println(slice)

	// 方式二
	var slice2 []string = []string{"1", "2", "1", "4"}
	fmt.Println(slice2)
	fmt.Println(cap(slice2))
	fmt.Println(len(slice2))

	// 遍历
	// for循环遍历
	for i := 0; i < len(slice); i++ {

		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}

	//for-range 遍历切片
	for i, v := range slice2 {

		fmt.Printf("i=%v v=%v \n", i, v)
	}
	/**
	  1. 切片初始化时，var slice = arr[StartIndex:endIndex]
	     说明：从arr数组下标为startIndex，取到下标为endIndex的元素（不包含arr[endIndex]）
	  2. 切片初始化时，仍然不能越界。范围在[0-len(arr)]之间，但是可以动态增长
	     1. var slice = arr[0:end],可以简写为 var slice = arr[:end]
	     2. var slice = arr[start:len(arr)] 可以简写：var slice = arr[start:]
	     3. var slice = arr[0:len(arr)] 可以简写：var slice = arr[:]
	  3. cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
	  4. 切片定义完之后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者make一个空间供切片来使用。
	  5. 切片可以继续使用切片

	*/

}
