package main

import (
	"fmt"
)

func main(){

	//continue break goto

label1:	// break label1 会跳到这里
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break label1
			}
			fmt.Printf("i=%v j=%v\n", i, j)
		}
	}

label2:	// break label1 会跳到这里
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue label2
			}
			fmt.Printf("i=%v j=%v\n", i, j)
		}
	}


	var n = 30
	if n > 24 {
		fmt.Println("成年人")
		goto label3
	}


	fmt.Println("aaa")
	fmt.Println("bbb")

label3:
	fmt.Println("ccc")
	fmt.Println("ddd")
}