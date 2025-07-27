package main

import "fmt"

func main() {

	//常量定义, 定义常量必须直接赋值，且值后续不能重新赋值
	const pi = 3.14159

	fmt.Println(pi)

	//多个常量可以一起声明, 如果省略了值则和上面的一行值一样

	const (
		A = "A"
		B = "B"
		C
	)

	fmt.Println(A, B, C)

	//iota 是 golang 语言的常量计数器，只能在常量的表达式中使用，能生成自增值

	const a = iota
	fmt.Println(a)

	//声明多个，值自增, _可跳过一个值
	const (
		n1 = iota
		n2
		n3
		n4
		_
		n5
	)

	fmt.Println(n1, n2, n3, n4, n5)

	//声明多个，值自增, _可跳过一个值
	const (
		b1 = iota //0
		b2 = 100  //100
		b3 = iota //2
		b4        //3

	)

	fmt.Println(b1, b2, b3, b4)

	//多个 iota 定义在一行
	const (
		x1, x2 = iota + 1, iota + 2
		x3, x4
		x5, x6
	)

	fmt.Println(x1, x2)
	fmt.Println(x3, x4)
	fmt.Println(x5, x6)

	//go fmt demo2.go 可以格式化代码

	

}
