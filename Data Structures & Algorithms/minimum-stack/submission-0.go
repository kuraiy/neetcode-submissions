type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	var mn int
	if len(this.minStack) == 0 {
		mn = val
	} else {
		mn = min(val, this.minStack[len(this.minStack)-1])
	}
	this.minStack = append(this.minStack, mn)
}

func (this *MinStack) Pop() {
	n := len(this.stack) - 1
	this.stack = this.stack[:n]
	this.minStack = this.minStack[:n]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}