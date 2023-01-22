package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// 接口
type File interface {
	GetContent() string
	GetNthLine(int) string
}

type FileEater struct {
	Path string
}

func (file FileEater) GetContent() string {
	f, err := ioutil.ReadFile(file.Path)
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f)
}

func (file FileEater) GetNthLine(n int) string {
	f, err := ioutil.ReadFile(file.Path)
	if err != nil {
		fmt.Println("read fail", err)
	}
	str := string(f)
	return strings.Split(str, "\n")[n-1]
}

func main() {
	file := FileEater{Path: "text.txt"}
	fmt.Print(file.GetContent())
	fmt.Print(file.GetNthLine(1))
	/*
		你好
		世界你好
	*/
}
