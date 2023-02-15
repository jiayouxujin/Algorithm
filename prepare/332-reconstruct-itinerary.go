package main

import "sort"

func findItinerary(tickets [][]string) []string {
	n := len(tickets)
	graph := make(map[string][]string)
	for i := 0; i < n; i++ {
		from, to := tickets[i][0], tickets[i][1]
		graph[from] = append(graph[from], to)
	}
	for key := range graph {
		sort.Strings(graph[key])
	}
	res := make([]string, 0)
	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			if len(graph[curr]) == 0 {
				break
			}
			next := graph[curr][0]
			graph[curr] = graph[curr][1:]
			dfs(next)
		}
		res = append(res, curr)
	}
	dfs("JFK")
	reverseArr(res)
	return res
}

func reverseArr(res []string) {
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
}
