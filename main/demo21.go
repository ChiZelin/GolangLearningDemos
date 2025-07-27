package main

import (
	"fmt"
)

//自定义类型
type myInt int

//类型别名
type myFloat = float64

//结构体， 结构体首字母大写的表示公有的  小写表示私有的  属性名大小写也同样是大写共有 小写私有
type Person struct {
	name string
	age int 
	sex string
}


func main(){

	var a myInt = 10
	fmt.Printf("%v %T\n", a, a) //10 main.myInt

	var b myFloat = 12.3
	fmt.Printf("%v %T\n", b, b) //12.3 float64 , 类型别名还是原类型的类型
	

	//结构体
	//实例化
	var p1 Person
	p1.name = "张三"
	p1.age = 20
	p1.sex = "男"
	fmt.Printf("值：%v 类型：%T\n", p1, p1) //值：{张三 20 男} 类型：main.Person
	//用 %#v 输出结构体信息  更详细
	fmt.Printf("值：%#v 类型：%T\n", p1, p1) //值：main.Person{name:"张三", age:20, sex:"男"} 类型：main.Person


	//用 new 创建结构体指针
	var p2 = new(Person)
	p2.name = "张三"
	p2.age = 20
	p2.sex = "男"

	(*p2).name = "王五" // 用指针 (*p2) 操作

	fmt.Printf("值：%v 类型：%T\n", p2, p2) //值：&{王五 20 男} 类型：*main.Person
	//用 %#v 输出结构体信息  更详细
	fmt.Printf("值：%#v 类型：%T\n", p2, p2) //值：&main.Person{name:"王五", age:20, sex:"男"} 类型：*main.Person

	//用 & 创建
	p3 := &Person{}
	p3.name = "张三"
	p3.age = 20
	p3.sex = "男"
	fmt.Printf("值：%v 类型：%T\n", p3, p3) //值：&{张三 20 男} 类型：*main.Person
	//用 %#v 输出结构体信息  更详细
	fmt.Printf("值：%#v 类型：%T\n", p3, p3) //值：&main.Person{name:"张三", age:20, sex:"男"} 类型：*main.Person


	//
	var p4 = Person{
		name:"哈哈",
		age:20,
		sex:"男",
	}

	fmt.Printf("值：%v 类型：%T\n", p4, p4) //值：{哈哈 20 男} 类型：main.Person
	//用 %#v 输出结构体信息  更详细
	fmt.Printf("值：%#v 类型：%T\n", p4, p4) //值：main.Person{name:"哈哈", age:20, sex:"男"} 类型：main.Person

}