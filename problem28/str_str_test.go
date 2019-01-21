package problem28

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	fmt.Printf("%s - %s - %d\n", "hello", "ll", strStr("hello", "ll"))
	fmt.Printf("%s - %s - %d\n", "aaaaa", "bba", strStr("aaaaa", "bba"))
	fmt.Printf("%s - %s - %d\n", "a", "a", strStr("a", "a"))
	fmt.Printf("%s - %s - %d\n", "", "a", strStr("", "a"))
	fmt.Printf("%s - %s - %d\n", "aaa", "aaa", strStr("aaa", "aaa"))
	fmt.Printf("%s - %s - %d\n", "aaa", "a", strStr("aaa", "a"))
}
