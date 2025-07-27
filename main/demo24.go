package main

import (
	"fmt"
)

//结构体嵌套

type Address struct {
	Phone string
	City  string
}

type Person struct {
	Name    string
	Age     int
	Hobby   []string
	map1    map[string]string
	Address Address //也可用匿名嵌套结构体 只写一个Address
}

func main() {

	//结构体嵌套

	var p Person
	p.Name = "张三"
	p.Age = 20
	p.Hobby = make([]string, 6, 6)
	p.Hobby[0] = "写代码"
	p.Hobby[1] = "打篮球"
	p.Hobby[2] = "睡觉"

	p.map1 = make(map[string]string)
	p.map1["K1"] = "V1"
	p.map1["K2"] = "V2"

	p.Address.Phone = "12345678" //连续. 给嵌套结构图赋值
	p.Address.City = "北京"

	fmt.Println(p)         //{张三 20 [写代码 打篮球 睡觉   ] map[K1:V1 K2:V2] {12345678 北京}}
	fmt.Printf("%#v\n", p) //main.Person{Name:"张三", Age:20, Hobby:[]string{"写代码", "打篮球", "睡觉", "", "", ""}, map1:map[string]string{"K1":"V1", "K2":"V2"}, Address:main.Address{Phone:"12345678", City:"北京"}}

	//在访问结构体中字段时，如果字段名与嵌套结构体中字段名一样
}
