package main

import (
	"encoding/json"
	"fmt"
)

type Sa struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func main() {

	// 结构体的使用细节

	var p3 *Sa = new(Sa)
	p3.Name = "123"
	p3.Age = 18
	fmt.Println(*p3)

	// 序列化成Json字符串
	//Marshal 使用到反射
	jsonString, _ := json.Marshal(p3)
	fmt.Println(string(jsonString))

}
