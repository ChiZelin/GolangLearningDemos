package main

import "fmt"

func getUserinfo() (string, int) {
	return "Bob", 20
}

func main (){

	// 1. var 声明变量
	//  var 变量名称 类型

	var username string

	fmt.Println(username) 


	var m_ = "李四"

	fmt.Println(m_)


	username = "张三"

	fmt.Println(username)


	//类型推导
	var username2 = "张三2"
	fmt.Println(username2)

	var age = 20
	var sex = "男"
	fmt.Println(username2, age, sex)


	//一次声明多个变量
	
	var a1, a2 string 

	a1 = "aaa"
	a2 = "aaaaaaa"
    fmt.Println(a1, a2)

	var (
		username3 string
		age3 int
		sex3 string
	)

	username3 = "张三"
	age3 = 20
	sex3 = "男"

	fmt.Println(username3, age3, sex3)

	var (
		username4 string = "张三4"
		age4 int = 20
		sex4 string = "男4"
	)

	fmt.Println(username4, age4, sex4)

	//短变量声明法, 短变量声明法只能声明局部变量，不能声明全局变量
	username5 := "张三5"
	age5 := 20
	sex5 := "男5"

	fmt.Println(username5, age5, sex5)


	//短变量声明法一次声明多个变量
	x, y, z := 1, 2, 3
	fmt.Println(x, y, z)

	x2, y2, z2 := 1, 2, "aaa"
	fmt.Println(x2, y2, z2)
	fmt.Printf("a类型%T   b类型%T   c类型%T\n", x2, y2, z2)


	var un, ag = getUserinfo()
	fmt.Println(un, ag)

	var un2,  _ = getUserinfo() //匿名变量，忽略一个值的时候用 _
	fmt.Println(un2)
}