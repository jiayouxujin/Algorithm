/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}

	if root.Left==nil{
		return minDepth(root.Right)+1
	}
	if root.Right==nil{
		return minDepth(root.Left)+1
	}

	return min(minDepth(root.Left),minDepth(root.Right))+1
}

func min(a,b int)int{
	if a<=b{
		return a
	}else{
		return b
	}
}
// @lc code=end

