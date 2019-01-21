package problem53

import (
	"fmt"
	"testing"
)

var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

func TestMaxSubArray(t *testing.T) {
	fmt.Printf("max sum: %d\n", maxSubArray(nums))
	fmt.Printf("max sum: %d\n", maxSubArray([]int{-1}))
	fmt.Printf("max sum: %d\n", maxSubArray([]int{-2, -1}))
	fmt.Printf("max sum: %d\n", maxSubArray([]int{-2, -1, -3, 3}))
	fmt.Printf("max sum: %d\n", maxSubArray([]int{1, -2, 0}))
}
