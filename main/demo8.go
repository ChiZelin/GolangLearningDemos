package main

import (
	"fmt"
)

func main(){

	//for range

	var str = "您好golang"

	for k, v := range str {
		fmt.Printf("key=%v v=%c\n", k, v)
	}


	//switch case
	var ext = "html"

	switch ext {
	case "html":
		fmt.Println("text/html")
		break
	case "css":
		fmt.Println("text/css")
		break
	default:
		fmt.Println("找不到此后缀")
	}

	//case可以有多个值, 变量定义也可以写到 switch case 里面, golang里面case里面不写break, 它不会继续执行下面的case, fallthrough可以让它向下穿透执行与它相邻的case
	//var n = 5
	switch n := 5; n {
	case 1, 3, 5:
		fmt.Println("奇数")
		break
		//fallthrough
	case 2, 4, 6:
		fmt.Println("偶数")
		break
	}

	//case 后面可以写表达式，这时 switch 语句后面就不需要再写变量了

	var age = 24

	switch {
	case age < 24:
		fmt.Println("好好学习")
	case age >= 24:
		fmt.Println("好好赚钱")	
	}


}