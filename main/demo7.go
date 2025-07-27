package main

import (
	"fmt"
)

func main(){

	//if 判断语句 if后面的大括号不能省略，且左括号得在 if 同行后面
	age := 30
	if age > 20 {
		fmt.Println("成年人")
	}

	// 可以在if 语句里面声明局部变量
	if age := 34; age > 20 {
		fmt.Println("成年人")
	}


	//for 循环
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//初始化部分放在循环外面
	j := 1
	for ; j <= 10; j++ {
		fmt.Println(j)
	}

	//for 循环可只写一个条件表达式
	k := 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}

	//无线循环, Golang 里面没有while，这种方式可以替代 while
	m := 1
	for{
		if m <= 10 {
			fmt.Println(m)
		} else {
			break
		}
		m++
	}

}