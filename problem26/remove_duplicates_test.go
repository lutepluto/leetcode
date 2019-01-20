package problem26

import (
	"fmt"
	"testing"
)

var nums1 = []int{1, 1, 2}
var nums2 = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

func TestRemoveDuplicates(t *testing.T) {
	fmt.Printf("Before removing: %v\n", nums1)
	fmt.Printf("After removing: %d - %v\n", removeDuplicates(nums1), nums1)

	fmt.Printf("Before removing: %v\n", nums2)
	fmt.Printf("After removing: %d - %v\n", solution(nums2), nums2)
}
