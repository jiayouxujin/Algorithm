package main

import (
	"strconv"
)

type TrieNode struct {
	Children [10]*TrieNode
}

func findKthNumber(n int, k int) int {
	if n < 1 || k < 1 {
		return 0
	}
	root := &TrieNode{}
	for i := 1; i <= n; i++ {
		curr := root
		for _, c := range strconv.Itoa(i) {
			index := int(c - '0')
			if curr.Children[index] == nil {
				curr.Children[index] = &TrieNode{}
			}
			curr = curr.Children[index]
		}
	}

	var res int
	queue := make([]int, 0)
	queue = append(queue, 1)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		k--
		if k == 0 {
			res = curr
			break
		}
		node := root.Children[curr]
		for i := 9; i >= 0; i-- {
			if node.Children[i] != nil {
				queue = append(queue, curr*10+i)
			}
		}
	}
	return res
}
