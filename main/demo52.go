package main

import (
	"fmt"
	"os"
)

func main() {
	// 删除空目录
	err := os.Remove("testdir")
	if err != nil {
		fmt.Println("删除空目录失败:", err)
	} else {
		fmt.Println("空目录删除成功")
	}

	// 递归删除目录（无论是否为空）
	err = os.RemoveAll("parent")
	if err != nil {
		fmt.Println("递归删除目录失败:", err)
	} else {
		fmt.Println("递归删除目录成功")
	}
}
