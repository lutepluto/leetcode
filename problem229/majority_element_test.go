package problem229

import (
	"fmt"
	"testing"
)

func TestMajorityElements(t *testing.T) {
	fmt.Printf("%#v\n", majorityElements([]int{3, 2, 3}))
	fmt.Printf("%#v\n", majorityElements([]int{1, 1, 1, 3, 3, 2, 2, 2}))
	fmt.Printf("%#v\n", majorityElements([]int{1, 2, 2, 3, 2, 1, 1, 3}))
}
