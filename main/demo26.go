package main

import (
	"fmt"
	"encoding/json"
)

type Student struct {
	Id int `json:"id"` //结构体标签
	Gender string
	Name string
	Sno string
}

func main(){

	//结构体和Json的转换
	
	var s1 = Student{
		Id: 12,
		Gender: "男",
		Name: "李四",
		Sno: "s0001",
	}

	fmt.Printf("%#v\n", s1) //main.Student{ID:12, Gender:"男", Name:"李四", Sno:"s0001"}

	//结构体转换为 json
	jsonByte, _ := json.Marshal(s1) //切片
	jsonStr := string(jsonByte)  //切片转为字符串
	fmt.Printf("%v\n", jsonStr)  //{"ID":12,"Gender":"男","Name":"李四","Sno":"s0001"}

	//json 转为 结构体
	var str = `{"ID":12,"Gender":"男","Name":"李四","Sno":"s0001"}`
	var s2 Student
	err := json.Unmarshal([]byte(str), &s2)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", s2) //main.Student{ID:12, Gender:"男", Name:"李四", Sno:"s0001"}
	
}