package main

func isValidBST(root *TreeNode) bool {
	return isValid2(root, nil, nil)
}

func isValid2(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isValid2(root.Left, min, root) && isValid2(root.Right, root, max)
}
