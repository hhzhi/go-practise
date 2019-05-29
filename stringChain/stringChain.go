package main

import (
	"fmt"
	"strings"
)

func main() {
	list := []string{
		"go a",
		"go b",
		"go c",
		"go d",
		"go e",
	}

	// 字符处理链
	chain := []func(string) string{
		delPrix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	StringParoccess(list, chain)

	for _, v := range list {
		fmt.Println(v)
	}
}

func delPrix(s string) string {
	return strings.TrimPrefix(s, "go")
}

func StringParoccess(list []string, chain []func(string) string) {

	for index, str := range list {

		item := str

		for _, f := range chain {
			item = f(item)
		}

		list[index] = item
	}
}
