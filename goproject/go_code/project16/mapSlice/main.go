package main

import "fmt"

func main() {

	// map 切片

	slice := make([]map[string]string, 2)

	if slice[0] == nil {

		slice[0] = make(map[string]string, 2)
		slice[0]["cities"] = "北京"
		slice[0]["age"] = "100"
	}
	// 动态增加
	newSlice := map[string]string{
		"name": "123",
		"agte": "123",
	}
	slice = append(slice, newSlice)

	fmt.Println(slice)

}
