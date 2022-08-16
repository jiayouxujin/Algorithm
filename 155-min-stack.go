package main

type MinStack struct {
	Data *[]int
	Min  *[]int
}

func Constructor3() MinStack {
	tmp1, tmp2 := []int{}, []int{}
	return MinStack{
		Data: &tmp1,
		Min:  &tmp2,
	}
}

func (this *MinStack) Push(val int) {
	if len(*this.Min) == 0 {
		*this.Min = append(*this.Min, val)
	} else {
		if val >= (*this.Min)[len(*this.Min)-1] {
			*this.Min = append(*this.Min, (*this.Min)[len(*this.Min)-1])
		} else {
			*this.Min = append(*this.Min, val)
		}
	}
	*this.Data = append(*this.Data, val)
}

func (this *MinStack) Pop() {
	*this.Data = (*this.Data)[:len(*this.Data)-1]
	*this.Min = (*this.Min)[:len(*this.Min)-1]
}

func (this *MinStack) Top() int {
	return (*this.Data)[len(*this.Data)-1]
}

func (this *MinStack) GetMin() int {
	return (*this.Min)[len(*this.Min)-1]
}
