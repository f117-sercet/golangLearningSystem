package main

import "fmt"

// 构造方式
type Person struct {
	Name string
	Age  int64
}

func main() {

	// 方式2：
	p1 := Person{"Marry", 20}

	fmt.Println(p1)

	// 方式三：
	var p3 *Person = new(Person)
	(*p3).Name = "tom"
	(*p3).Age = 18
	fmt.Println(*p3)

	// 方式3简化

	var p4 *Person = new(Person)
	p4.Name = "toms"
	(*p3).Age = 19
	fmt.Println("方式三简化", p4)

	// 方式4
	var person *Person = &Person{}

	person.Name = "王二"
	person.Age = 26

	fmt.Print("方式4", *person)
}
