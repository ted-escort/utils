package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

// CreateDir 创建目录
func CreateDir(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false, nil
		}
	}
	return true, nil
}

// FileIsExist 文件是否存在
func FileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// ReadFile 读文件
func ReadFile(file string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(file)
	// 若出现错误，打印错误提示
	if err != nil {
		fmt.Println("解析配置文件错误：", err.Error())
		return nil, err
	}
	return bytes, nil
}
