package main

//结构体实现多个接口

import (
	"fmt"
)

type Usber1 interface {
	start()
}

type Usber2 interface {
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

func main() {

	p := Phone{
		Name: "华为手机",
	}

	p.start()
	p.stop()

	var p1 Usber1
	p1 = p

	var p2 Usber2
	p2 = p

	p1.start()
	p2.stop()

}
