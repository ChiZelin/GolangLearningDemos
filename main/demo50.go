package main

import (
	"fmt"
	"io/ioutil"
)

//文件读写
//3种写方法
//1.file.write()  2.bufio  3.ioutil

func main() {
	//3.ioutil
	str := "hello golang, cool!"
	err := ioutil.WriteFile("./test.txt", []byte(str), 0666)

	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("写入成功")
}
