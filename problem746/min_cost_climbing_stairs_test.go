package problem746

import (
	"fmt"
	"testing"
)

var cost1 = []int{10, 15, 20}
var cost2 = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}

func TestMinCostClimbingStairs(t *testing.T) {
	fmt.Printf("Cheapest cost on %#v is %d\n", cost1, minCostClimbingStairs(cost1))
	fmt.Printf("Cheapest cost on %#v is %d\n", cost2, minCostClimbingStairs(cost2))
}
