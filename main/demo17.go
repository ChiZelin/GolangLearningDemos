package main

import (
	"fmt"
)

func fn1() {
	fmt.Println("fn1")
}

func fn2() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	panic("抛出一个异常")
}

func main(){
	//panic recover
	fn1()
	fn2()
	fmt.Println("结束")
	

	
}