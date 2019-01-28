package problem70

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	fmt.Printf("There are %d ways to climb to %d staires.\n", climbStairs(2), 2)
	fmt.Printf("There are %d ways to climb to %d staires.\n", climbStairs(3), 3)
}
