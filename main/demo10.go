package main

import (
	"fmt"
)

func main(){

	//数组, 数组的长度是数组的一部分

	var arr1 [3]int
	var arr2 [4]int
	var strArr [3]string

	fmt.Printf("arr1:%T  arr2:%T strArr:%T", arr1, arr2, strArr) //arr1:[3]int  arr2:[4]int strArr:[3]string  数组的长度是数组的一部分

	//初始化

	arr1[0] = 23
	arr1[1] = 21
	arr1[2] = 25

	fmt.Println(arr1)


	var arr3 = [4]int{23, 24, 5, 3}
	fmt.Println(arr3)

	var arr4 = [...]int{23, 24, 5, 3} // ... 可省略数组的长度
	fmt.Println(arr4)


	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for k, v := range arr1 {
		fmt.Printf("k:%v  v:%v\n", k, v)
	}


}