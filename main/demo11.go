package main

import (
	"fmt"
)

func main(){

	//值类型  引用类型  基础数据类型 和 数组都是值类型

	var a = 10
	b := a 
	a = 20

	fmt.Println(a, b)

	var arr1 = [...]int{1, 2, 3}

	arr2 := arr1
	arr1[0] = 11

	fmt.Println(arr1, arr2) //[11 2 3] [1 2 3]


	//切片是引用数据类型

	var arr3 = []int {1, 2, 3} //[]中不写... 不写数字   是切片类型
	arr4 := arr3

	arr3[0]=11

	fmt.Println(arr3, arr4) //[11 2 3] [11 2 3]


	//二维数组，第一个[]中可以用...来推断行数
	var arr5 = [3][2]string{
		{"a","b"},
		{"c","d"},
		{"e","f"},
	}

	fmt.Println(arr5[0][1])

	//遍历二维数组
	for _, v := range arr5 {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
}