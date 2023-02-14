package main

func sumNumbers(root *TreeNode) int {
	return dfs2(root, 0)
}

func dfs2(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs2(root.Left, sum) + dfs2(root.Right, sum)
}
