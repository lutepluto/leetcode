package problem136

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	fmt.Printf("%v\n", singleNumber([]int{2, 2, 1}))
	fmt.Printf("%v\n", singleNumber([]int{4, 1, 2, 1, 2}))
}
