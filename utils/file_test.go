package utils

import (
	"fmt"
	"testing"
)

func Test_GetFileList(t *testing.T) {
	path := "."
	fileList, err := GetFileList(path, true)
	if err != nil {
		fmt.Println("获取文件列表失败")
	} else {
		fmt.Println(fileList)
	}
}

func Test_CopyFile(t *testing.T) {
	srcFile := "file_test.go"
	distFile := "file_test.out"
	size, err := CopyFile(srcFile, distFile)
	if err != nil {
		fmt.Println("复制文件失败")
	} else {
		fmt.Printf("成功复制了 %d 字节", size)
	}
}

func Test_CopyDir(t *testing.T) {
	srcDir := "."
	distDir := "./../test.out"
	err := CopyDir(srcDir, distDir)
	if err != nil {
		fmt.Println("复制目录失败")
	} else {
		fmt.Printf("复制目录成功")
	}
}
