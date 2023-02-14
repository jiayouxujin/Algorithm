package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var res []string
	dfs3(s, 0, []string{}, &res)
	return res
}

func dfs3(s string, index int, path []string, res *[]string) {
	if index == 4 {
		if len(s) == 0 {
			*res = append(*res, strings.Join(path, "."))
		}
		return
	}
	for i := 1; i <= 3; i++ {
		if i > len(s) {
			break
		}
		if i > 1 && s[0] == '0' {
			break
		}
		num, _ := strconv.Atoi(s[:i])
		if num > 255 {
			break
		}
		dfs3(s[i:], index+1, append(path, s[:i]), res)
	}
}
