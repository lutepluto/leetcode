package problem118

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	triangle := generate(5)
	for _, row := range triangle {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Printf("\n")
	}
}
