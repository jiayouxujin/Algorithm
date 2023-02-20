package main

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue, index, maxWidth := make([]*TreeNode, 0), make([]int, 0), 0
	queue = append(queue, root)
	index = append(index, 0)
	for len(queue) > 0 {
		size, left, right := len(queue), 0, 0
		for i := 0; i < size; i++ {
			node, idx := queue[0], index[0]
			queue, index = queue[1:], index[1:]
			if i == 0 {
				left = idx
			}
			if i == size-1 {
				right = idx
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
				index = append(index, 2*idx)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				index = append(index, 2*idx+1)
			}
		}
		maxWidth = max(maxWidth, right-left+1)
	}
	return maxWidth
}
