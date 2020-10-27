/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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
func pathSum(root *TreeNode, sum int) [][]int {
	var slice [][]int
	slice=findPath(root,sum,slice,make([]int,0))
	return slice
}

func findPath(root *TreeNode,sum int,slice [][]int,stack []int)[][]int{
	if root==nil{
		return slice
	}

	sum-=root.Val
	stack=append(stack,root.Val)

	if sum==0&&root.Left==nil&&root.Right==nil{
		slice=append(slice,append([]int{},stack...))
		stack=stack[:len(stack)-1]
	}

	slice=findPath(root.Left,sum,slice,stack)
	slice=findPath(root.Right,sum,slice,stack)

	return slice
}
// @lc code=end

