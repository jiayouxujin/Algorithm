package main

type CQueue struct {
	in, out []int
}

func Constructor3() CQueue {
	return CQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 {
		for i := len(this.in) - 1; i >= 0; i-- {
			v := this.in[i]
			this.in = this.in[:len(this.in)-1]
			this.out = append(this.out, v)
		}
	}
	v := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return v
}
