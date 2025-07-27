package main

import (
	"fmt"
	"reflect"
)

//反射

// 通过反射获取变量类型
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("类型：%v     类型名称：%v    类型种类：%v \n", v, v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	//fmt.Println(v)

	// var n = v.Int() + 12 // 可以通过v.Int() 来获取 int的值
	// fmt.Println(n)

	kind := v.Kind()

	switch kind {
	case reflect.Int:
		fmt.Printf("int 类型的原始值%v\n", v.Int())
	case reflect.Float32:
		fmt.Printf("Float32 类型的原始值%f\n", v.Float())
	case reflect.Float64:
		fmt.Printf("Float64 类型的原始值%v\n", v.Float())
	case reflect.String:
		fmt.Printf("String 类型的原始值%v\n", v.String())
	default:
		fmt.Printf("还没有判断这个类型\n")

	}

}

func reflectSetValue(x interface{}) {
	// v, _ := x.(*int64)
	// *v = 123

	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())        //ptr
	fmt.Println(v.Elem().Kind()) //int64  通过 .Elem() 获取指针锁指向的实际数据的数据类型

	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(123) //通过 .Elem() 获取指针锁指向的实际数据的数据类型
	}

}

type myInt int

type Person struct {
	Name string
	Age  int
}

func main() {

	a := 10
	b := 23.4
	c := true
	d := "你好golang"

	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)

	var e myInt = 34
	var f = Person{
		Name: "张三",
		Age:  20,
	}

	reflectFn(e)
	reflectFn(f)

	var h = 25
	reflectFn(&h) //类型Name为空 Kind为 ptr

	//数组
	var i = [3]int{1, 2, 3} //类型Name为空 Kind为 array
	reflectFn(i)

	//切片
	var j = []int{1, 2, 3} //类型Name为空 Kind为 slice
	reflectFn(j)

	/*output
	类型：int     类型名称：int    类型种类：int
	类型：float64     类型名称：float64    类型种类：float64
	类型：bool     类型名称：bool    类型种类：bool
	类型：string     类型名称：string    类型种类：string
	类型：main.myInt     类型名称：myInt    类型种类：int
	类型：main.Person     类型名称：Person    类型种类：struct
	类型：*int     类型名称：    类型种类：ptr
	类型：[3]int     类型名称：    类型种类：array
	类型：[]int     类型名称：    类型种类：slice
	*/

	//通过反射获取值
	var k = 13
	reflectValue(k)
	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
	reflectValue(d)

	//通过反射修改值
	var n int64 = 100
	fmt.Println(n)
	reflectSetValue(&n)
	fmt.Println(n)

}
