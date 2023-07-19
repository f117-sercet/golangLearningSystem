package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {

	rType := reflect.TypeOf(b)
	rValue := reflect.ValueOf(b)
	i := rValue.Int()
	fmt.Println(rType.Kind())
	fmt.Println(rValue)
	fmt.Println(i)

	iv := rValue.Interface()
	num2 := iv.(int)
	fmt.Println(num2)

}

func main() {
	of := reflect.ValueOf(1)

	fmt.Println(of)

	var num int = 100

	reflectTest(num)

}
