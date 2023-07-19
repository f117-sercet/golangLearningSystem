package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflectTest(b interface{}) {

	rType := reflect.TypeOf(b)
	rValue := reflect.ValueOf(b)
	fmt.Println(rType)
	fmt.Println(rValue)

	iv := rValue.Interface()
	student, ok := iv.(Student)

	if ok {
		fmt.Println(student.Name)
		fmt.Println(ok)
	}
}

func main() {

	stu := Student{
		Name: "tom",
		Age:  20,
	}

	reflectTest(stu)

}
