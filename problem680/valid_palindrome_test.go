package problem680

import (
	"fmt"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	fmt.Printf("%v\n", validPalindrome("abc"))
	fmt.Printf("%v\n", validPalindrome("aba"))
	fmt.Printf("%v\n", validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))
}
