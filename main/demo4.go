package main

import (
	"fmt"
)

func main(){

	//bool 类型, 不能强制类型转换，默认值 false
	var flag = true

	fmt.Printf("%v--%T\n", flag, flag)


	if flag {
		fmt.Printf("True")
	}
	
}