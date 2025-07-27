package main

import (
	"fmt"
)

func main(){

	//map, 也是引用数据类型
	//make 创建 map
	var userinfo = make(map[string]string)

	userinfo["username"] = "张三"
	userinfo["age"] = "20"
	userinfo["sex"] = "男"
	fmt.Println(userinfo)
	fmt.Println(userinfo["username"])

	var userinfo2 = map[string]string{
		"username" : "李四",
		"age" : "23",
		"sex" : "男",
	}

	fmt.Println(userinfo2)
	fmt.Println(userinfo2["username"])

	//遍历
	for k, v := range userinfo {
		fmt.Printf("K:%v v:%v\n", k, v)
	}

	//判断key是否存在 ok 是 true 则存在
	v, ok := userinfo["username"]
	fmt.Printf("v:%v ok:%v\n", v, ok)

	//删除
	delete(userinfo, "sex")
	fmt.Println(userinfo)

	//map 类型的切片
	var userinfoSlice = make([]map[string]string, 3, 3)

	if userinfoSlice[0] == nil {
		userinfoSlice[0] = make(map[string]string)
		userinfoSlice[0]["username"] = "张三"
		userinfoSlice[0]["age"] = "20"
	}

	if userinfoSlice[1] == nil {
		userinfoSlice[1] = make(map[string]string)
		userinfoSlice[1]["username"] = "李四"
		userinfoSlice[1]["age"] = "23"
	}


	fmt.Println(userinfoSlice)


	//map 类型的值定义为一个切片
	var userinfo3 = make(map[string][]string)

	userinfo3["hobby"] = []string{
		"读书",
		"敲代码",
	}

	fmt.Println(userinfo3) //map[hobby:[读书 敲代码]]

}