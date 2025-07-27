package main

import "fmt"

//空接口 表示没有任何约束  任意的类型都可以实现空接口

type A interface{}

func show(a interface{}) { //空接口可以表示任意类型 可以当作任意类型来使用
	fmt.Printf("值：%v 类型：%T\n", a, a)
}

func PrintType(x interface{}) {
	if _, ok := x.(string); ok {
		fmt.Println("string 类型")
	} else if _, ok := x.(int); ok {
		fmt.Println("int 类型")
	}
}

func main() {

	var a A
	var str = "你好 golang"

	a = str
	fmt.Printf("值：%v 类型：%T\n", a, a)

	show("aaaa")
	show(203)
	show(true)
	slice := []int{1, 2, 3, 4}
	show(slice)

	var m1 = make(map[string]interface{}) //map的值类型通过空接口定义为任意类型
	m1["username"] = "张三"
	m1["age"] = 20

	fmt.Printf("值：%v 类型：%T\n", m1, m1)
	fmt.Printf("值：%v 类型：%T\n", m1["username"], m1["username"])
	fmt.Printf("值：%v 类型：%T\n", m1["age"], m1["age"])

	//类型断言
	var n interface{}

	n = "Hello golang"
	v, ok := n.(string) // 用 x.(T) 来断言 x 是不是 T 类型

	if ok {
		fmt.Println("n就是一个string类型 值是:", v) //n就是一个string类型 值是: Hello golang
	}

	PrintType(33)
	PrintType("aaaa")

	// 如果结构体方法接收者是值类型 那么结构体对象的值类型和指针类型都可以赋值个接口变量
	// 如果结构体方法接收者是指针类型 那么结构体对象的值类型不能赋值给接口 指针类型能赋值给接口

	//以空接口接收的值有时需要结合断言来读取 如下例子所示
	var userinfo = make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["hobby"] = []string{"睡觉", "吃饭"}

	fmt.Println(userinfo["username"])
	fmt.Println(userinfo["hobby"])
	//fmt.Println(userinfo["hobby"][0]) //invalid operation: cannot index userinfo["hobby"] (map index expression of type interface{})

	//利用断言类获取切片
	hobby, _ := userinfo["hobby"].([]string)
	fmt.Println(hobby[0])

}
