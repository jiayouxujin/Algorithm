package main

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q, reachEnd := make([]*TreeNode, 0), false
	q = append(q, root)
	for len(q) > 0 {
		p := q
		q = make([]*TreeNode, 0)
		for i := 0; i < len(p); i++ {
			if p[i] == nil {
				reachEnd = true
				continue
			}
			if p[i] != nil && reachEnd == true {
				return false
			}
			q = append(q, p[i].Left)
			q = append(q, p[i].Right)
		}
	}
	return true
}
