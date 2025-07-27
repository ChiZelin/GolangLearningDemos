package main

import (
	"fmt"
)

//结构体的匿名字段，声明是只有类型 没有字段
type Person struct {
	string
	int
}

func main(){

	//结构体的匿名字段

	p := Person{
		"张三",
		20,
	}

	fmt.Println(p) //{张三 20}

	
}