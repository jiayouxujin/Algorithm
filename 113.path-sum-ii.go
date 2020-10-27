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
    if root==nil{
		return [][]int{}
	}

	if root.Left==nil&&root.Right==nil{
		if sum==root.Val{
			return [][]int{[]int{root.Val}}
		}
	}

	path,res:=[]int{},[][]int{}
	tmpLeft:=pathSum(root.Left,sum-root.Val)
	path=append(path,root.Val)
	if len(tmpLeft)>0{
		for i:=0;i<len(tmpLeft);i++{
			tmpLeft[i]=append(path,tmpLeft[i]...)
		}
		res=append(res,tmpLeft...)
	}

	path=[]int{}
	tmpRight:=pathSum(root.Right,sum-root.Val)
	path=append(path,root.Val)

	if len(tmpRight)>0{
		for i:=0;i<len(tmpRight);i++{
			tmpRight[i]=append(path,tmpRight[i]...)
		}
		res=append(res,tmpRight...)
	}

	return res;
}
// @lc code=end

