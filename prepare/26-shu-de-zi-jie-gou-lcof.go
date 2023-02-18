package main

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return dfsSubStructure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func dfsSubStructure(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return dfsSubStructure(A.Left, B.Left) && dfsSubStructure(A.Right, B.Right)
}
