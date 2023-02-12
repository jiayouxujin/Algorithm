package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	var nodes []*TreeNode
	if root == nil {
		return nil
	}
	nodes = append(nodes, root)

	for len(nodes) > 0 {
		var tmp []*TreeNode
		var cur []int
		for _, node := range nodes {
			cur = append(cur, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		nodes = tmp
		if len(res)%2 == 0 {
			res = append(res, cur)
		} else {
			res = append(res, reverse(cur))
		}
	}
	return res
}

func reverse(arr []int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}
