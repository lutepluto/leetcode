package problem168

import (
	"fmt"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	fmt.Printf("%s\n", convertToTitle(1))
	fmt.Printf("%s\n", convertToTitle(28))
	fmt.Printf("%s\n", convertToTitle(701))
	fmt.Printf("%s\n", convertToTitle(702))
}
