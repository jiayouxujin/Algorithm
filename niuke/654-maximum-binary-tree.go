package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func constructMaximumBinaryTree(nums []int) *TreeNode {
//	return constructTree(nums, 0, len(nums)-1)
//}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	stack := make([]*TreeNode, 0)
	for _, num := range nums {
		node := &TreeNode{Val: num}
		for len(stack) > 0 && stack[len(stack)-1].Val < num {
			node.Left = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			stack[len(stack)-1].Right = node
		}
		stack = append(stack, node)
	}
	return stack[0]
}

func constructTree(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	if left == right {
		return &TreeNode{Val: nums[left]}
	}
	maxIndex := left
	for i := left; i <= right; i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	node := &TreeNode{Val: nums[maxIndex]}
	node.Left = constructTree(nums, left, maxIndex-1)
	node.Right = constructTree(nums, maxIndex+1, right)
	return node
}
