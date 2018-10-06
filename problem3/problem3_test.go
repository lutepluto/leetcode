package problem3

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Printf("abcabcbb: %d\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("bbbbb: %d\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Printf("pwwkew: %d\n", lengthOfLongestSubstring("pwwkew"))
	fmt.Printf("abba: %d\n", lengthOfLongestSubstring("abba"))
}
