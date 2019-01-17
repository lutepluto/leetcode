package problem14

import (
	"fmt"
	"testing"
)

var strs1 = []string{"flower", "flow", "flight"}
var strs2 = []string{"dog", "racecar", "car"}

func TestHorizontalScanning(t *testing.T) {
	fmt.Printf("longest common prefix of %v is: %s\n", strs1, horizontalScanning(strs1))
	fmt.Printf("longest common prefix of %v is: %s\n", strs2, horizontalScanning(strs2))
}

func TestVerticalScannig(t *testing.T) {
	fmt.Printf("longest common prefix of %v is: %s\n", strs1, verticalScanning(strs1))
	fmt.Printf("longest common prefix of %v is: %s\n", strs2, verticalScanning(strs2))
}

func TestDivideAndConquer(t *testing.T) {
	fmt.Printf("longest common prefix of %v is: %s\n", strs1, divideAndConquer(strs1))
	fmt.Printf("longest common prefix of %v is: %s\n", strs2, divideAndConquer(strs2))
}
