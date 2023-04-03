package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// 获取目录下的文件列表
// dir 获取的目录
// recursion 是否递归查找子目录
func GetFileList(dir string, recursion bool) (res []string, err error) {
	dir = strings.TrimRight(dir, "/")
	fs, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, f := range fs {
		fullName := path.Join(dir, f.Name())
		if f.IsDir() {
			if recursion {
				var subRes []string
				subRes, err = GetFileList(fullName, true)
				if err != nil {
					return
				}
				res = append(res, subRes...)
			}
			continue
		}
		res = append(res, fullName)
	}
	return
}

// 复制文件
// src 复制源文件
// dist 目标文件
func CopyFile(src, dist string) (copySize int, err error) {
	// 打开源文件
	sFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer sFile.Close()

	// 打开写入文件
	dFile, err := os.Create(dist)
	if err != nil {
		return
	}
	defer dFile.Close()

	// 创建缓冲区，边读边写
	buf := make([]byte, 4096) // 由于虚拟内存的最小单位是 page(默认是4096)，所以，设置成 4096
	readSize := 0
	for {
		readSize, err = sFile.Read(buf)
		if err != nil && err != io.EOF { // 读取错误
			return
		}
		if readSize == 0 { // 读完毕
			err = nil
			return
		}
		copySize += readSize
		// 写入文件
		dFile.Write(buf[:readSize])
	}
}

// 复制目录
// src 复制源目录
// dist 目标目录
func CopyDir(src, dist string) error {
	finfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !finfo.IsDir() {
		return fmt.Errorf("%s is not a directory", src)
	}
	fileInfos, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	// 确保目标目录存在
	if info, err := os.Stat(dist); err != nil || (info != nil && !info.IsDir()) {
		if err = os.Mkdir(dist, os.ModePerm); err != nil {
			return err
		}
	}

	// 遍历目录项，递归进行赋值
	for _, fileInfo := range fileInfos {
		srcFile := path.Join(src, fileInfo.Name())
		distFile := path.Join(dist, fileInfo.Name())
		if fileInfo.IsDir() {
			if err = CopyDir(srcFile, distFile); err != nil {
				return err
			}
		} else if _, err = CopyFile(path.Join(src, fileInfo.Name()), path.Join(dist, fileInfo.Name())); err != nil {
			return err
		}
	}
	return nil
}
