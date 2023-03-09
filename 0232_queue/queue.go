package queue

type MyQueue struct {
	input  []int
	output []int
}

func Constructor() MyQueue {
	return MyQueue{
		input:  make([]int, 0, 4),
		output: make([]int, 0, 4),
	}
}

func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}

	this.input2output()
	val := this.output[len(this.output)-1]
	this.output = this.output[:len(this.output)-1]
	return val
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}

	this.input2output()
	return this.output[len(this.output)-1]
}

func (this *MyQueue) input2output() {
	if len(this.output) != 0 {
		return
	}

	for i := len(this.input) - 1; i >= 0; i-- {
		this.output = append(this.output, this.input[i])
	}
	this.input = make([]int, 0, 4)
}

func (this *MyQueue) Empty() bool {
	return len(this.output) == 0 && len(this.input) == 0
}
