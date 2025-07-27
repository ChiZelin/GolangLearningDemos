package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//文件读写
//3种读方法
//1.file.Read()  2.bufio  3.ioutil

func main() {
	//2.bufio

	//只读方式打开当前目录下的 main.go文件   D:/go-code/project1/main/hello.go
	file, err := os.Open("./hello.go")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	//bufio 读取文件
	var fileStr string

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n') // \n 表示一次读取一行

		if err == io.EOF {
			fileStr += str //注意这里读取完毕 str中可能还会包含一些数据
			fmt.Println("读取完毕")
			break
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		fileStr += str
	}

	fmt.Println(fileStr)

}
