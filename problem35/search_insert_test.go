package problem35

import (
	"fmt"
	"testing"
)

var nums = []int{1, 3, 5, 6}

func TestSearchInsert(t *testing.T) {
	fmt.Printf("%#v inert %d at %d\n", nums, 5, searchInsert(nums, 5))
	fmt.Printf("%#v inert %d at %d\n", nums, 2, searchInsert(nums, 2))
	fmt.Printf("%#v inert %d at %d\n", nums, 7, searchInsert(nums, 7))
	fmt.Printf("%#v inert %d at %d\n", nums, 0, searchInsert(nums, 0))
}
