package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSortDesc(t *testing.T) {
	// Init
	elements := []int{9, 4, 17, 3, 2, 0, 8}
	// Execution
	fmt.Println(elements)
	BubbleSortDesc(elements)
	fmt.Println(elements)
	// Validation
	if elements[0] != 17 {
		t.Error("First element should be 17")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("Last element should be 0")
	}
}
