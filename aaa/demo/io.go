package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
Reader: 读取器
*/

//	os.O_WRONLY 只写
//	os.O_CREATE 创建文件
//	os.RDONLY  只读
//	os.RDWR  读写
//	os.TRUNC 清空
//	os.O_APPEND 追加
func WriteFile(s string) {
	// w写 r读 x执行 w:2 r:4
	file, err := os.OpenFile("./auth.txt", os.O_CREATE|os.O_WRONLY, 0666)
	fmt.Println("write", s)
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(s)
	if err != nil {
		fmt.Println("write file error", err)
	}
}

func ReadFile() (string, error) {
	var filePath = "./auth.txt"
	_, err := os.Stat(filePath)
	if err != nil {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("create file error", file)
			return "", err
		}
	}
	file, err := os.Open("./auth.txt")

	if err != nil {
		fmt.Println("open file error: ", err)
		return "", err
	}
	defer file.Close()
	// 定义接收文件读取的字节数组
	//var buf [128]byte
	var content []byte
	reader := bufio.NewReader(file)
	for {
		// 按行读取
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file error", err)
			return "", err
		}
		//content = append(content, buf[:n]...)
		content = line
	}
	return string(content), nil
}
