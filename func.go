package main

import "strings"

// 以行为单位，翻转输出
func ReverseOutput(f File) string {
	lines := len(strings.Split(f.GetContent(), "\n"))
	var out string
	for i := lines; i >= 1; i-- {
		out += f.GetNthLine(i) + "\n"
	}
	return out
}
