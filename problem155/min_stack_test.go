package problem155

import (
	"fmt"
	"testing"
)

func TestGetMin(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Printf("%v\n", minStack.GetMin())
	minStack.Pop()
	minStack.Pop()
	fmt.Printf("%v\n", minStack.GetMin())
}
