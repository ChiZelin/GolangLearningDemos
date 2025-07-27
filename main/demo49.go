package main

import (
	"bufio"
	"fmt"
	"os"
)

//文件读写
//3种写方法
//1.file.write()  2.bufio  3.ioutil

func main() {
	//2.bufio
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666) //os.O_APPEND
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	writer := bufio.NewWriter(file)
	writer.WriteString("你好golang") //将数据先写入缓存
	writer.Flush()

	fmt.Println("写入成功")
}
