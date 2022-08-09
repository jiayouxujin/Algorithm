package main

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	maxDepth2(root, &res)
	return res
}

func maxDepth2(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth2(root.Left, res)
	rightMax := maxDepth2(root.Right, res)

	*res = max(*res, leftMax+rightMax)
	return 1 + max(leftMax, rightMax)
}
