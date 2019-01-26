package problem66

import (
	"fmt"
	"testing"
)

var digits1 = []int{1, 2, 3}
var digits2 = []int{4, 3, 2, 1}
var digits3 = []int{9}
var digits4 = []int{8, 3, 4, 9}

func TestPlusOne(t *testing.T) {
	fmt.Printf("%#v plus one equals to %#v\n", digits1, plusOne(digits1))
	fmt.Printf("%#v plus one equals to %#v\n", digits2, plusOne(digits2))
	fmt.Printf("%#v plus one equals to %#v\n", digits3, plusOne(digits3))
	fmt.Printf("%#v plus one equals to %#v\n", digits4, plusOne(digits4))
}
