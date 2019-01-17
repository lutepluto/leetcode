package problem13

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	fmt.Printf("%s is %d\n", "III", romanToInt("III"))
	fmt.Printf("%s is %d\n", "IV", romanToInt("IV"))
	fmt.Printf("%s is %d\n", "IX", romanToInt("IX"))
	fmt.Printf("%s is %d\n", "LVIII", romanToInt("LVIII"))
	fmt.Printf("%s is %d\n", "MCMXCIV", romanToInt("MCMXCIV"))
}
