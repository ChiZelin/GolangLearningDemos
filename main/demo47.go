package main

import (
	"fmt"
	"io/ioutil"
)

//文件读写
//3种读方法
//1.file.Read()  2.bufio  3.ioutil

func main() {
	//3.ioutil
	//文件比较小时用这个， 但注意这个方法已经 deprecated 了

	byteStr, err := ioutil.ReadFile("./hello.go")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(byteStr))

}
