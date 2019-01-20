package problem27

import (
	"fmt"
	"testing"
)

var nums1 = []int{2, 3, 3}
var nums2 = []int{0, 1, 2, 2, 3, 0, 4, 2}
var nums3 = []int{5}

func TestRemoveElement(t *testing.T) {
	fmt.Printf("Before removing: %v\n", nums1)
	fmt.Printf("After removing: %d - %v\n", removeElement(nums1, 3), nums1)

	fmt.Printf("Before removing: %v\n", nums2)
	fmt.Printf("After removing: %d - %v\n", removeElement(nums2, 2), nums2)

	fmt.Printf("Before removing: %v\n", nums3)
	fmt.Printf("After removing: %d - %v\n", removeElement(nums3, 5), nums3)
}
