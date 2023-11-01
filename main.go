package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// 获取 Finder 当前路径
	getPath()
	exitIterm2App()

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

func exitIterm2App() bool {
	appBundleID := "com.googlecode.iterm2"
	cmd := exec.Command("mdfind", fmt.Sprintf("kMDItemCFBundleIdentifier == '%s'", appBundleID))
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("未安装应用程序\n")
	} else {
		if len(strings.TrimSpace(string(output))) > 0 {
			fmt.Printf("已安装应用程序\n")
			return true
		} else {
			fmt.Printf("未安装应用程序\n")
		}
	}
	return false

}
