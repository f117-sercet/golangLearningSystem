package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {

	rVal := reflect.ValueOf(b)
	fmt.Println("rVal的Kind", rVal.Kind())

	rVal.Elem().SetInt(4)

}

func main() {

	var num int = 2

	reflectTest(&num)
	fmt.Println("num", num)

}
