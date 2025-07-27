package main

import (
	"fmt"
)

//注意事项： 非本地类型不能定义方法，也就是说不能给别的包里面的类型定义方法

type Person struct {
	name string
	age int 
	sex string
}

func (p Person) PrintInfo(){
	fmt.Printf("姓名： %v 年龄：%v\n", p.name, p.age)
}

//接收者是指针类型 此时才能修改值 因为是用指针 修改同一块内存 而不是修改另一块内存
func (p *Person) SetInfo(name string, age int){
	p.name = name
	p.age = age
}


//还可以给自定义类型定义方法
type myInt int
func (m myInt) SayHello(){
	fmt.Println("Hello, 我是一个int")
}


func main(){

	//结构体方法 和 接收者

	var p1 = Person{
		name:"张三",
		age:20,
		sex:"男",
	}

	p1.PrintInfo()

	var p2 = Person{
		name:"李四",
		age:20,
		sex:"男",
	}

	p2.PrintInfo()

	p1.SetInfo("王五", 20)

	p1.PrintInfo()

	var mi myInt = 10

	mi.SayHello()
}