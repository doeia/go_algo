package main

import "fmt"

type Usb interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "start")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "stop")
}

func main() {
	p := Phone{
		Name: "iphone",
	}
	// p.start()

	var p1 Usb
	p1 = p
	p1.start()

	// 空接口

	// 断言
}
