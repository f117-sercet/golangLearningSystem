package main

import "fmt"

type Boy struct {
	Id   int
	Next *Boy
}

// 环形链表-约瑟夫问题

func AddBoy(num int) *Boy {

	first := &Boy{}
	curBoy := &Boy{}
	//boy3 := Boy{}
	if num < 1 {
		return first
	}

	for i := 1; i <= num; i++ {
		if i == 1 {
			boy := &Boy{
				Id: i,
			}
			if i == 1 {
				first = boy
				curBoy = boy
				curBoy.Next = first
			} else {
				curBoy.Next = boy
				curBoy = boy
				curBoy.Next = boy
			}
		}
	}
	return first
}
func ShowBoy(first *Boy) {
	if first.Next == nil {
		return
	}
	curBoy := first
	for {

		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}
func PlayGame(first *Boy, startNo int, countNum int) {

	// 空链表
	if first.Next == nil {

		fmt.Println("链表为空")

		return
	}
	tail := first
	for {

		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	for i := 0; i < startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	for {
		for i := 0; i < countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈 \n\n\n", first.Id)
		first = first.Next
		tail.Next = first

		if first == tail {
			break
		}
	}
	fmt.Printf("小孩编号为%d 出圈 \n\n", first.Id)
}

func main() {

	first := AddBoy(50)
	ShowBoy(first)
	PlayGame(first, 2, 3)
}
