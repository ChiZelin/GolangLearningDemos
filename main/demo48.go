package main

import (
	"fmt"
	"os"
)

//文件读写
//3种写方法
//1.file.write()  2.bufio  3.ioutil

func main() {
	//1.file.write()

	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666) //os.O_APPEND
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	// WriteString
	// for i := 0; i < 10; i++ {
	// 	//写入文件
	// 	file.WriteString("直接写入字符串数据" + strconv.Itoa(i) + "\r\n") //记事本里面是\r\n识别换行
	// }

	//file.write
	var str = "直接写入字符串"
	file.Write([]byte(str))

	fmt.Println("写入成功")
}
