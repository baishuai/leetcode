package p155

type MinStack struct {
	stack  []int
	min    []int
	length int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: make([]int, 0), min: make([]int, 0), length: 0}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if this.length > 0 && this.min[this.length-1] < x {
		x = this.min[this.length-1]
	}
	this.min = append(this.min, x)
	this.length++
}

func (this *MinStack) Pop() {
	if this.length == 0 {
		return
	}
	this.stack = this.stack[:this.length-1]
	this.min = this.min[:this.length-1]
	this.length--
}

func (this *MinStack) Top() int {
	return this.stack[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.length-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
