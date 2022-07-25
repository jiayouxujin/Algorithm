package main

func sumNumbers(root *TreeNode) int {
	res := 0
	dfsSum(root, 0, &res)
	return res
}

func dfsSum(root *TreeNode, sum int, res *int) {
	if root == nil {
		return
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += sum
	}
	dfsSum(root.Left, sum, res)
	dfsSum(root.Right, sum, res)
}
