package main

import "fmt"

func fn1(n int) int {
	if n > 0 {
		
		return n + fn1(n-1)
	} else {
		return n
	}
}


//闭包定义
//闭包让变量常驻内存 不污染全局
func adder() func() int{
	var i = 10

	return func() int {
		return i + 1
	}
}

func adder2() func(y int) int{
	var i = 10

	return func(y int) int {
		i += y
		return i
	}
}

func main() {

	//递归调用
	fmt.Println(fn1(10))


	//闭包
	var fn = adder()

	fmt.Println(fn()) //11
	fmt.Println(fn()) //11
	fmt.Println(fn()) //11

	var fn2 = adder2()

	fmt.Println(fn2(1)) //11
	fmt.Println(fn2(1)) //12
	fmt.Println(fn2(1)) //13


}
