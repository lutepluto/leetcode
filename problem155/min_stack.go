package problem155

// MinStack is a stack which tracks the minimum values
type MinStack struct {
	stack []int
	mins  []int
}

// Constructor returns a new MinStack instance
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		mins:  make([]int, 0),
	}
}

// Push pushes value into the MinStack
func (minSstack *MinStack) Push(x int) {
	minSstack.stack = append(minSstack.stack, x)
	minsLen := len(minSstack.mins)
	if minsLen == 0 || minSstack.mins[minsLen-1] >= x {
		minSstack.mins = append(minSstack.mins, x)
	}
}

// Pop removes the last value pushed into the MinStack and returns it
func (minSstack *MinStack) Pop() {
	top := minSstack.Top()
	minSstack.stack = minSstack.stack[:len(minSstack.stack)-1]

	minsLen := len(minSstack.mins)
	if minsLen > 0 && minSstack.mins[minsLen-1] == top {
		minSstack.mins = minSstack.mins[:minsLen-1]
	}
}

// Top returns the last value pushed into the MinStack
func (minSstack *MinStack) Top() int {
	len := len(minSstack.stack)
	return minSstack.stack[len-1]
}

// GetMin returns the minimum value in the MinStack
func (minSstack *MinStack) GetMin() int {
	return minSstack.mins[len(minSstack.mins)-1]
}
