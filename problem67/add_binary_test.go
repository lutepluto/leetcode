package problem67

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	fmt.Printf("%s + %s = %s\n", "11", "1", addBinary("11", "1"))
	fmt.Printf("%s + %s = %s\n", "1010", "1011", addBinary("1010", "1011"))
}
