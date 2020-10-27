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
	var res [][]int
	res=findPath(root,sum,res,make([]int,0))
	return res;
}

func findPath(root *TreeNode,sum int,res [][]int,cur []int)[][]int{
	if root==nil{
		return res
	}

	sum-=root.Val
	cur=append(cur,root.Val)

	if sum==0&&root.Left==nil&&root.Right==nil{
		res=append(res,append([]int{},cur...))
		cur=append(cur,len(cur)-1)
	}

	res=findPath(root.Left,sum,res,cur)
	res=findPath(root.Right,sum,res,cur)
	return res
}
// @lc code=end

