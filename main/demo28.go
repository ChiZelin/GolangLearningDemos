package main

import "fmt"

//接口

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

type Camera struct {
}

func (p Camera) start() {
	fmt.Println("相机启动")
}

func (p Camera) stop() {
	fmt.Println("相机关机")
}

type Computer struct {
}

func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
}

func main() {

	p := Phone{
		Name: "华为手机",
	}

	p.start()
	p.stop()

	var p1 Usber
	p1 = p //表示 phone 实现了 Usber 接口

	p1.start()
	p1.stop()

	c := Camera{}

	var c1 Usber = c
	c1.start()
	c1.stop()
	fmt.Printf("类型%T\n", c1) //类型main.Camera

	var computer = Computer{}
	var phone = Phone{
		Name: "小米",
	}
	var camera = Camera{}

	computer.work(phone)
	computer.work(camera)

}
