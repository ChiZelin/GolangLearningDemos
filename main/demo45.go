package main

import (
	"fmt"
	"io"
	"os"
)

//文件读写
//3种读方法
//1.file.Read()  2.bufio  3.ioutil

func main() {

	//1.file.Read()

	//只读方式打开当前目录下的 main.go文件   D:/go-code/project1/main/hello.go
	file, err := os.Open("./hello.go")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(file)

	//读取文件里面的内容
	var strSlice []byte
	var tempSlice = make([]byte, 128)

	for {
		n, err := file.Read(tempSlice)

		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}

		if err != nil {
			fmt.Println("读取失败")
			return
		}

		strSlice = append(strSlice, tempSlice[:n]...) //注意截取到所读的位置即可
	}

	//fmt.Printf("读取到了%v个字节\n", n)
	fmt.Println(string(strSlice))

}
