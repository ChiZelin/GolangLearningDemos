package main

import (
	"fmt"
)


func f1() {
	fmt.Println("开始")
	defer func() {
		fmt.Println("aaaa")
	}()
	fmt.Println("结束")
}

func f2() int {
	var a int
	defer func() {
		a++
	}()
	
	return a 
}

func f3() (a int) {

	defer func() {
		a++
	}()
	
	return a 
}

func f4() (a int) {

	defer func() {
		a++
	}()
	
	return 5 
}

func main(){

	//defer
	// fmt.Println("开始")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("结束")
	//输出 开始 结束 3 2 1

	a2 := f2()

	fmt.Println(a2) //0

	a3 := f3()

	fmt.Println(a3) //1

	a4 := f4()

	fmt.Println(a4) //6

}