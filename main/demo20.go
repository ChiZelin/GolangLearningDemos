package main

import (
	"fmt"
)

func main(){

	//指针
	var a int = 10
	var p = &a
	fmt.Printf("a的值： %v a的类型%T  a的地址%p\n", a, a, &a) //a的值： 10 a的类型int  a的地址0xc00000a0a8
	fmt.Printf("a的值： %v a的类型%T  a的地址%p\n", p, p, &p) //a的值： 0xc00000a0a8 a的类型*int  a的地址0xc000052020

	fmt.Println(*p) //*p表示取出所指向地址的值

	*p = 30

	fmt.Println(*p)
	fmt.Println(a)


	//new 给指针类型分配存储空间， 实际项目中用的不是很多
	var a2 = new(int) //指针类型 类型是 *int  指向的值0

	fmt.Printf("a2的值： %v a2的类型%T  a2指向的地址的值%v\n", a2, a2, *a2)
	

	//make 用于slice map 以及 channel 的初始化
}