package main

import "fmt"

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(YB)
}
