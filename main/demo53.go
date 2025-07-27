package main

import (
	"fmt"
	"os"
)

func main() {
	// 原文件路径
	oldPath := "oldname.txt"
	// 新文件路径（可包含目录）
	newPath := "newname.txt"

	// 创建示例文件（实际使用时可跳过此步骤）
	f, _ := os.Create(oldPath)
	f.Close()

	// 执行重命名操作
	err := os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println("重命名失败:", err)
		return
	}
	fmt.Println("文件重命名成功")
}
