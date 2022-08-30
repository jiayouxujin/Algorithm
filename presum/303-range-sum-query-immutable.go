package presum

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArray{
		preSum: preSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}
