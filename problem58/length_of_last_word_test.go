package problem58

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	fmt.Printf("%s's last word has the length of %d\n", "Hello World", lengthOfLastWord("Hello World"))
	fmt.Printf("%s's last word has the length of %d\n", "  ", lengthOfLastWord("  "))
	fmt.Printf("%s's last word has the length of %d\n", "leetcode coding  ", lengthOfLastWord("leetcode coding  "))
}
