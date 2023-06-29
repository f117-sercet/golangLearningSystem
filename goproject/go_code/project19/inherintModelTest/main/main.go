package main

import "fmt"

// 继承
type Goods struct {
	Name  string
	price int
}

func (goods *Goods) Get() *Goods {

	return &Goods{
		Name:  "123",
		price: 33,
	}

}

type Books struct {
	//匿名结构体
	Goods
	writer string
}

/*
func Get(goods *Goods) {

	fmt.Println("123")

}
*/
func main() {

	books := new(Books)

	books.Goods.price = 24
	books.Name = "360"
	books.writer = "3"
	goods := books.Goods.Get()
	fmt.Print(goods)
	fmt.Println(books)

}
