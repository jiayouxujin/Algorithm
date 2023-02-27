package main

import "strconv"

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
}
