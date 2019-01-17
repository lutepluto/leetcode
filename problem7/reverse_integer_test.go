package problem7

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Printf("%d's reverse int is %d\n", 123, reverse(123))
	fmt.Printf("%d's reverse int is %d\n", -123, reverse(-123))
	fmt.Printf("%d's reverse int is %d\n", 120, reverse(120))
}
