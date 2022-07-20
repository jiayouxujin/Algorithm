package main

import "math"

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := math.MinInt32
	getPathSum(root, &maxSum)
	return maxSum
}

func getPathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return math.MinInt32
	}
	left := getPathSum(root.Left, maxSum)
	right := getPathSum(root.Right, maxSum)

	curMax := max(max(left+root.Val, right+root.Val), root.Val)
	*maxSum = max(*maxSum, max(curMax, left+right+root.Val))
	return curMax
}
