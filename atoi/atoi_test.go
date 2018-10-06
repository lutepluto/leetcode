package atoi

import (
	"fmt"
	"testing"
)

func TestAtoi(t *testing.T) {
	fmt.Printf("%s is %d\n", "  ", atoi("  "))
	fmt.Printf("%s is %d\n", " +", atoi(" +"))
	fmt.Printf("%s is %d\n", "42", atoi("42"))
	fmt.Printf("%s is %d\n", "-42", atoi("-42"))
	fmt.Printf("%s is %d\n", "4193 with words", atoi("4193 with words"))
	fmt.Printf("%s is %d\n", "words and 987", atoi("words and 987"))
	fmt.Printf("%s is %d\n", "-91283472332", atoi("-91283472332"))

	// str := "+12adwi中概股"
	// for _, runeValue := range str {
	// 	fmt.Printf("%v %v\n", runeValue, string(runeValue))
	// }
}
