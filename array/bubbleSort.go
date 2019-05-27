package main

import "fmt"

func main() {
	a := []int{2, 1, 35, 32, 52, 6}

	fmt.Println(a)

	num := len(a)
	for i := 0; i < num-1; i++ {
		for j := 0; j < num-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	fmt.Println(a)
}
