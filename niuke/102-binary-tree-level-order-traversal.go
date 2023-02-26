package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res, queue := make([][]int, 0), make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		siz, tmp := len(queue), make([]int, 0)
		for i := 0; i < siz; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}
