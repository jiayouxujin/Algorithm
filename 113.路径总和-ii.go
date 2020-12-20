/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
	var res [][]int
	return helper(root,sum,res,[]int(nil))
}

func helper(root *TreeNode,sum int,res [][]int,tmp []int)[][]int{
	if root==nil{
		return res
	}
	sum-=root.Val
	tmp=append(tmp,root.Val)
	if root.Left==nil&&root.Right==nil&&sum==0{
		res=append(res,append([]int{},tmp...))
		tmp=tmp[:len(tmp)-1]
	}

	res=helper(root.Left,sum,res,tmp)
	res=helper(root.Right,sum,res,tmp)
	return res
}
// @lc code=end

