package main

import (
	"fmt"
	"time"
)

func fibonacci() func() int {
	a, b := 0, 1
	sum := 0
	return func() int {
		if sum == 0 {
			sum++
			return a
		}
		a, b = b, (a + b)
		return a
	}
}

func main() {
	f := fibonacci()

	start := time.Now()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	fmt.Println(time.Since(start))
}
