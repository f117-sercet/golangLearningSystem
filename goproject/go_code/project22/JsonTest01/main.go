package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func testMap() {
	var a map[string]string
	a = make(map[string]string)
	a["name"] = "123"
	a["age"] = "123"
	marshal, _ := json.Marshal(a)
	fmt.Println("map序列化后的结结果", string(marshal))
}
func testSlice() {

	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "123"
	m1["token"] = "33"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})

	m2["name"] = "444"
	m2["token"] = "55"
	slice = append(slice, m2)

	marshal, _ := json.Marshal(slice)
	fmt.Println("切片序列化", string(marshal))

}

func main() {

	/*	var s string = "123"

		marshal, err := json.Marshal(s)
		fmt.Println(string(marshal), err)*/

	// 序列化
	var student Student
	student.Name = "123"
	student.Age = 18
	data, _ := json.Marshal(student)
	fmt.Println(string(data))

	testMap()
	testSlice()

}
