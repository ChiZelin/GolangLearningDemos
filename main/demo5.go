package main

import (
	"fmt"
	"strings"
)

func main(){

	//字符串

	var str1 string = "你好 golang"

	fmt.Printf("%v--%T\n", str1, str1)

	//字符串转义

	str2  := "this is str \\ "
	fmt.Println(str2)

	//多行字符串
	str3 := `
	this is a string
	this is a string2
	`
	fmt.Println(str3)

	//len(str) 求长度
	var str4 = "你好"
	fmt.Println(len(str4))  //1个汉字占用3个字节

	//+ 或者 fmt.Sprintf 拼接字符串
	str5 := "你好"
	str6 := "Golang"

	fmt.Println(str5 + str6)
	str7 := fmt.Sprintf("%v %v",str5, str6) //Sprintf 拼接之后能赋值
	fmt.Println(str7)

	//也可以用 + 拼，但 + 不能在每行开头
	str8 := "123" + 
	"456" +
	"789"
	fmt.Println(str8)


	var str9 = "123-456-789"
	arr := strings.Split(str9, "-")
	fmt.Println(arr) //[123 456 789]   切片


	str10 := strings.Join(arr, "*") //join 把切片拼成字符串
	fmt.Println(str10) //123*456*789

	arr2 := []string{"php", "java", "Golang"} //切片
	fmt.Println(arr2)


	// Contains
	str11 := "this is a str"
	str12 := "this"
	fmt.Println(strings.Contains(str11, str12)) 

	// strings.HasPrefix, strings.HasSuffix
	fmt.Println(strings.HasPrefix(str11, str12)) 
	fmt.Println(strings.HasSuffix(str11, str12)) 


	//strings.Index(), strings.LastIndex(), 字串出现的位置， 查找不到返回 -1
	str13 := "this is a string"
	str14 := "is"
	fmt.Println(strings.Index(str13, str14)) //2
	fmt.Println(strings.LastIndex(str13, str14)) //5

	
}