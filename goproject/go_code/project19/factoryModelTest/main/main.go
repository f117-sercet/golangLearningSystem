package main

import (
	"fmt"
	"golangLearningSystem/goproject/go_code/project19/factoryModelTest/model"
)

func main() {

	stu := model.Student{
		Name: "tom",
		Age:  18,
	}
	fmt.Println(stu)
	fmt.Println(123)

	//工厂模式
	student := model.NewStudent("小王", 18)
	fmt.Println(*student)

}
