package main

func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res, stack := make([]int, 0), make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}
