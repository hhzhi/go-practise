package bubblesort

import (
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 2, 3, 4, 1}

	BubbleSort(values)
	if values[0] != 1 {
		t.Error("BubbleSort() failed. Got", values, "Expected 12345")
	}
}
