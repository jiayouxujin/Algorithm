package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	return helper(root, targetSum, res, []int{})
}

func helper(root *TreeNode, sum int, res [][]int, tmp []int) [][]int {
	if root == nil {
		return res
	}
	sum -= root.Val
	tmp = append(tmp, root.Val)
	if root.Left == nil && root.Right == nil && sum == 0 {
		res = append(res, append([]int{}, tmp...))
		tmp = tmp[:len(tmp)-1]
	}

	res = helper(root.Left, sum, res, tmp)
	res = helper(root.Right, sum, res, tmp)
	return res
}
