package sorting

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	fmt.Printf("%v\n", bubbleSort([]int{4, 2, 9, 6, 1, 5}))
	fmt.Printf("%v\n", selectionSort([]int{4, 2, 9, 6, 1, 5}))
	fmt.Printf("%v\n", insertionSort([]int{7, 2, 3, 1, 8, 4, 10}))
	fmt.Printf("%v\n", quickSort([]int{7, 2, 3, 1, 8, 4, 10}))
}
