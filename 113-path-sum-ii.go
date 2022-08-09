package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	helper(root, targetSum, &res, []int{})
	return res
}

func helper(root *TreeNode, sum int, res *[][]int, tmp []int) {
	if root == nil {
		return
	}
	sum -= root.Val

	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			tmp = append(tmp, root.Val)
			t := make([]int, len(tmp))
			copy(t, tmp)
			*res = append(*res, t)
		}
		return
	}

	tmp = append(tmp, root.Val)
	helper(root.Left, sum, res, tmp)
	tmp = tmp[:len(tmp)-1]

	tmp = append(tmp, root.Val)
	helper(root.Right, sum, res, tmp)
	tmp = tmp[:len(tmp)-1]
}
