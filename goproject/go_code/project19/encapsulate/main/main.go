package main

import (
	"fmt"
	"golangLearningSystem/goproject/go_code/project19/encapsulate/model"
)

func main() {

	// 封装实例
	p := model.NewPerson("smith")
	p.SetSal(4500)
	p.SetAge(28)

	fmt.Println(*p)

}
