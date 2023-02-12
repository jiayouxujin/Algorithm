package main

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) in2out() {
	for len(this.in) > 0 {
		this.out = append(this.out, this.in[len(this.in)-1])
		this.in = this.in[:len(this.in)-1]
	}
}

func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		this.in2out()
	}
	val := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		this.in2out()
	}
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in)+len(this.out) == 0
}
