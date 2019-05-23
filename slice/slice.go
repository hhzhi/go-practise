package main

import (
	"fmt"
)

func main() {
	var mySlice = make([]int, 5, 6)

	mySlice = []int{1, 2, 3, 4, 5, 6}
	slice1 := mySlice[3:]
	slice2 := mySlice[:3]
	slice3 := []int{45, 46}

	// cap
	fmt.Println("cap(slice1) = ", cap(slice1))

	// append
	slice1 = append(slice1, slice2...)
	slice1 = append(slice1, 11, 22)

	// copy
	copy(slice1, slice3)
	copy(slice3, slice2)

	for _, v := range slice1 {
		fmt.Println(" v = ", v)
	}

	for i := 0; i < len(slice3); i++ {
		fmt.Println("slice3 item ", i, " = ", slice3[i])
	}
}
