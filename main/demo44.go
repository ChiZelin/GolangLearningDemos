package main

import (
	"fmt"
	"reflect"
)

//结构体反射

type Student struct {
	Name  string `json:"name" form:"username"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名：%v 年龄：%v 成绩：%v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Println("这是一个打印方法...")
}

func PrintStructField(s interface{}) {

	//判断是不是结构体
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}

	fmt.Println("传入的参数是一个结构体")

	//1.通过类型变量里面的Field可以获取结构体的字段

	field0 := t.Field(0) //这种方式可以用在遍历的时候
	fmt.Printf("%#v", field0)
	fmt.Printf("字段名称:%v\n", field0.Name)             //通过field获取字段名
	fmt.Printf("字段类型:%v\n", field0.Type)             //通过field获取字段类型
	fmt.Printf("字段Tag:%v\n", field0.Tag.Get("json")) //通过field获取字段的 Tag 中的 json 的值
	fmt.Printf("字段Tag:%v\n", field0.Tag.Get("form")) //通过field获取字段的 Tag 中的 form 的值

	fmt.Println("---------------------------")

	//2.通过类型变量里面的FieldByName可以获取结构体的字段
	field1, ok := t.FieldByName("Age")

	if ok {
		fmt.Printf("字段名称:%v\n", field1.Name)             //通过field获取字段名
		fmt.Printf("字段类型:%v\n", field1.Type)             //通过field获取字段类型
		fmt.Printf("字段Tag:%v\n", field1.Tag.Get("json")) //通过field获取字段的 Tag 中的 json 的值
	}

	//3.通过类型变量里面的NumField获取到该结构体有几个字段
	var fieldCount = t.NumField()
	fmt.Println("结构体有", fieldCount, "个属性")

	//4.通过值变量获取结构体属性对应的值
	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))
	fmt.Println("---------------------------")

	for i := 0; i < fieldCount; i++ {
		fmt.Printf("属性名称:%v 属性值：%v 属性类型：%v 属性Tag:%v\n", t.Field(i).Name, v.Field(i), t.Field(i).Type, t.Field(i).Tag.Get("json"))
	}

}

// 打印执行方法
func PrintStructFn(s interface{}) {
	//1.通过类型变量里面的Method 可以获取结构体方法
	//判断是不是结构体
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}

	method0 := t.Method(0)    //和结构体方法的顺序没有关系，和结构体方法的ASCII有关系
	fmt.Println(method0.Name) //GetInfo
	fmt.Println(method0.Type) //func(main.Student) string

	//2.通过类型变量获取这个结构体有多少个方法
	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(method1.Name)
		fmt.Println(method1.Type)
	}

	//3.通过 值类型 获取方法 并调用
	v.Method(1).Call(nil)
	v.MethodByName("Print").Call(nil) //通过方法名获取方法

	info1 := v.MethodByName("GetInfo").Call(nil)

	fmt.Println(info1)

	//4.执行方法 传参数  需要注意 传递参数是 []reflect.Value的切片
	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(23))
	params = append(params, reflect.ValueOf(99))
	v.MethodByName("SetInfo").Call(params) //执行方法 传入参数

	info2 := v.MethodByName("GetInfo").Call(nil)

	fmt.Println(info2)

	//5.获取方法数量

	fmt.Println("方法数量：", t.NumMethod())

}

// 反射修改结构体属性
func reflectChangeStruct(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体指针类型")
		return
	}

	//修改结构体属性的值
	name := v.Elem().FieldByName("Name")
	name.SetString("小李")

	age := v.Elem().FieldByName("Age")
	age.SetInt(22)

}

func main() {

	stu1 := Student{
		Name:  "小明",
		Age:   15,
		Score: 98,
	}

	PrintStructField(stu1)
	PrintStructFn(&stu1)
	reflectChangeStruct(&stu1)

	fmt.Printf("%#v\n", stu1)
}
