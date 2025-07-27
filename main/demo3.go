package main

import(
	"fmt"
	"unsafe"
)

func main() {
	//基本数据类型
	//int int8 int16 int32 int64
	var num int8 = 127

	fmt.Printf("num=%v 类型：%T", num, num)

	var a int8 = -15

	fmt.Printf("num=%v 类型：%T\n", a, a)
	fmt.Println(unsafe.Sizeof(a)) //输出a占用的字节

	var n1 uint8 = 2
	fmt.Printf("num=%v 类型：%T\n", n1, n1)

	var a1 int32 = 10
	var a2 int64 = 21

	fmt.Println(int64(a1) + a2) //用 int64 进行强制类型转换

	//高位向低位转换,注意溢出了的问题 这个转完变 -126 了
	var n2 int16 = 130
	fmt.Println(int8(n2))

	num1 := 30
	fmt.Printf("num=%v 类型：%T\n", num1, num1)
	fmt.Println(unsafe.Sizeof(num1)) //输出a占用的字节
	fmt.Printf("num=%d\n", num1) //10进制输出
	fmt.Printf("num=%b\n", num1) //2进制输出
	fmt.Printf("num=%o\n", num1) //8进制输出
	fmt.Printf("num=%x\n", num1) //16进制输出


	//浮点数 float32, float64
	var f float32 = 3.12
	fmt.Printf("值：%v--%f,类型%T\n",f, f, f)
	fmt.Println(unsafe.Sizeof(f))

	var f2 float64 = 3.1287857456736573
	fmt.Printf("值：%v--%f--%.2f,类型%T\n",f2, f2, f2, f2)
	fmt.Println(unsafe.Sizeof(f2))

	var f3 = 3.14e2  // 科学计数法

	fmt.Printf("值：%v--%f--%.2f,类型%T\n",f3, f3, f3, f3)


	// float 数据精度丢失的问题

	var f4 float64 = 1129.6
	fmt.Println(f4 * 100)   // 输出 112959.99999999999 ， 解决办法使用第三方包 decimal 

	// int 转 float
	i := 10
	f5 := float64(i)
	
	fmt.Printf("值：%v", f5)


	// float 转 int
	f6 := 23.6
	i2 := int(f6)
	fmt.Printf("值：%v", i2)
}