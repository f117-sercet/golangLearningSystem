package main

import "fmt"

// 声明/顶一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}
type Computer struct {
}

// 让Phone 实现usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

// 只要是实现了USB接口，指实现了usb接口声明的所有方法
func (c Computer) Working(usb Usb) {
	usb.Stop()
	usb.Start()
}

func main() {
	var phone Phone
	phone.Start()
	phone.Stop()

	var computer Computer
	computer.Working(phone)
	computer.Working(phone)

	//fmt.Println(Phones())

}
