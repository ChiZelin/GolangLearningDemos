package main

import (
	"fmt"
)

//结构体继承

type Animal struct {
	Name string
}

func (a Animal) run(){
	fmt.Printf("%v 在运动\n", a.Name)
}

type Dog struct {
	Age int
	Animal //结构体嵌套 继承  这地方也可以嵌套指针 *Animal
}

func (d Dog) wang() {
	fmt.Printf("%v 在旺旺\n", d.Name)
}

func main(){

	//结构体嵌套 表示继承

	var d = Dog{
		Age: 20,
		Animal: Animal{  //如果嵌套的是指针 这地放 赋值地址 &Animal
			Name: "阿奇",
		},
	}

	d.run()
	d.wang()

}