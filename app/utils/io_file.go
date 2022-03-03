package utils

import (
	"fmt"
	"io/ioutil"
)

// ReadFile 文件读取
func ReadFile() (string, error) {
	content, err := ioutil.ReadFile("./auth.txt")
	if err != nil {
		fmt.Println("read file error---->>", err)
		return "", err
	}
	return string(content), err
}

// WriteFile 对文件写入
func WriteFile(s string) {
	err := ioutil.WriteFile("./auth.txt", []byte(s), 0666)
	if err != nil {
		fmt.Println("write file error ---->", err)
		return
	}
}
