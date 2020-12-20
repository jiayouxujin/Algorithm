/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
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
func sumNumbers(root *TreeNode) int {
	var res int
	helper(root,&res,0)
	return res
}

func helper(root *TreeNode,res *int,cur int){
	if root==nil{
		return 
	}
	cur=cur*10+root.Val
	if root.Left==nil&&root.Right==nil{
		*res+=cur
		cur=(cur-root.Val)/10
		return 
	}
	helper(root.Left,res,cur)
	helper(root.Right,res,cur)
	return 
}
// @lc code=end

