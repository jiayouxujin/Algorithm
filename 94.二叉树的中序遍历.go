package algorithm
/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 
 func inorderTraversal(root *TreeNode) []int {
	var res []int
	helper(root,&res)
	return res
}

func helper(root *TreeNode,res *[]int){
	if root==nil{
		return 
	}
	helper(root.Left,res)
	*res=append(*res,root.Val)
	helper(root.Right,res)
	
}
// @lc code=end

