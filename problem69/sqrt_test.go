package problem69

import (
	"fmt"
	"testing"
)

func TestSqrt(t *testing.T) {
	fmt.Printf("the square root of %d is %d\n", 4, sqrt(4))
	fmt.Printf("the square root of %d is %d\n", 8, sqrt(8))
	fmt.Printf("the square root of %d is %d\n", 9, sqrt(9))
}
