package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建单个目录（需要父目录存在）
	err := os.Mkdir("testdir", 0755)
	if err != nil {
		fmt.Println("创建目录失败:", err)
	} else {
		fmt.Println("目录创建成功")
	}

	// 创建多级目录（自动创建不存在的父目录）
	err = os.MkdirAll("parent/child/grandchild", 0755)
	if err != nil {
		fmt.Println("创建多级目录失败:", err)
	} else {
		fmt.Println("多级目录创建成功")
	}
}
