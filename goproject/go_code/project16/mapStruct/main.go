package main

import (
	"fmt"
)

func main() {

	// map 的声明
	var a map[string]string
	//fmt.Println(a)
	//使用map前，需要先make，作用是给map分配数据空间
	a = make(map[string]string, 20)
	a["number"] = "123"
	fmt.Printf("a=%p\n", &a)

	/***
	map的声明方式
	*/

	// 方式一:
	// map 的声明
	var a1 map[string]string
	//fmt.Println(a)
	//使用map前，需要先make，作用是给map分配数据空间
	a1 = make(map[string]string, 20)
	a1["number"] = "123"
	fmt.Printf("a=%p\n", &a1)

	// 方式二：
	cities := make(map[string]string)
	cities["2"] = "北京"
	cities["3"] = "天津"
	cities["4"] = "上海"
	cities["5"] = "福州"
	cities["6"] = "河南"
	fmt.Println(cities)

	// 方式三:
	b := map[string]string{

		"name":  "1",
		"names": "2",
		"named": "3",
	}
	fmt.Println(b)

	// 删除
	delete(b, "name")
	fmt.Println(b)

	val, ok := b["names"]

	if ok {
		fmt.Printf("有值 值为 %v\n\n", val)
	} else {
		fmt.Printf("没有 %v\n\n", val)
	}
}
