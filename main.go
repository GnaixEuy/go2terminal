package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	path := getPath()
	// 判断 操作系统
	switch os := runtime.GOOS; os {
	case "windows":
		windowsOpen(path)
	case "darwin":
		macOSOpen(path)
	case "linux":
		linuxOpen(path)
	default:
		fmt.Println("其他操作系统")
	}
}

func getPath() string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前路径时出错:", err)
		return ""
	}
	fmt.Println("当前路径:", currentDir)
	return currentDir
}
