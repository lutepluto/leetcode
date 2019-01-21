package problem38

import (
	"fmt"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	fmt.Printf("%d - %s\n", 1, countAndSay(1))
	fmt.Printf("%d - %s\n", 2, countAndSay(2))
	fmt.Printf("%d - %s\n", 3, countAndSay(3))
	fmt.Printf("%d - %s\n", 4, countAndSay(4))
	fmt.Printf("%d - %s\n", 5, countAndSay(5))
	fmt.Printf("%d - %s\n", 6, countAndSay(6))
	fmt.Printf("%d - %s\n", 7, countAndSay(7))
}
