package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	res, count := 0, 0
	traverse(root, k, &res, &count)
	return res
}

func traverse(root *TreeNode, k int, res, count *int) {
	if root != nil {
		traverse(root.Left, k, res, count)
		*count++
		if *count == k {
			*res = root.Val
			return
		}
		traverse(root.Right, k, res, count)
	}
}
