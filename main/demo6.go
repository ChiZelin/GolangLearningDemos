package main

import (
	"fmt"
	_"strings"
	_"unsafe"
	"strconv"
)

func main(){

	//字符

	var a = 'a'
	fmt.Printf("ASCII：%v 字符%c 类型：%T\n", a, a, a) //ASCII：97 字符a 类型：int32

	var str string = "this"
	fmt.Printf("ASCII：%v 字符%c 类型：%T\n", str[2], str[2], str[2]) //ASCII：97 字符a 类型：int32


	//unsafe.Sizeof() 没法查看string 类型的数据所占的存储空间  可以用len
	//fmt.Println(unsafe.Sizeof(str)) 
	fmt.Println(len(str))


	//byte 和 rune
	str2 := "你好，世界！"
    // 按字节遍历
    for i := 0; i < len(str2); i++ {
        fmt.Printf("字节 %d: %x\n", i, str2[i])
    }
    // 按rune遍历
    for i, r := range str2 {
        fmt.Printf("字符 %d: %c (%U)\n", i, r, r)
    }


	//通过 strconv 把其他类型转为 string
	var i int = 20
	str3 := strconv.FormatInt(int64(i), 10)
	fmt.Printf("值：%v 类型：%T\n", str3, str3)


	//运算符
	var a2 = 10
	var b2 = 3
	fmt.Println(a2 * b2)
	fmt.Println(a2 + b2)
	fmt.Println(a2 - b2)
	fmt.Println(a2 / b2)
	fmt.Println(a2 % b2)

	//++ -- 单独使用 不能赋值
	var aa = 10
	aa++
	fmt.Println(aa)
	aa--
	fmt.Println(aa)

	// == != > >= < <= && || ! += -= *= /= %=

	//位运算符  & | ^ << >>

	var a3 = 5 //101
	var b3 = 2 //010
	fmt.Println("a3&b3=",a3&b3) //000
	fmt.Println("a3|b3=",a3|b3) //111
	fmt.Println("a3^b3=",a3^b3) //111
	fmt.Println("a3<<b3=",a3<<b3) //100
	fmt.Println("a3>>b3=",a3>>b3) //001


}