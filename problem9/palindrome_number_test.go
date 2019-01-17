package problem9

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Printf("%d: %t\n", 121, isPalindrome(121))
	fmt.Printf("%d: %t\n", -121, isPalindrome(-121))
	fmt.Printf("%d: %t\n", 10, isPalindrome(10))
}
