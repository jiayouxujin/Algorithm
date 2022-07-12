package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		sz := len(queue)
		tmp := make([]int, 0)

		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			tmp = append(tmp, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		if len(res)%2 != 0 {
			reverseArray(tmp)
		}
		res = append(res, tmp)
	}
	return res
}

func reverseArray(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
