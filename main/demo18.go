package main

import (
	"fmt"
)

func fn1(a int, b int) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error:", err)
		}
	}()

	fmt.Println(a/b)
}

func main(){
	//panic recover
	
	fn1(10, 0)
	fmt.Println("结束")
	

	
}