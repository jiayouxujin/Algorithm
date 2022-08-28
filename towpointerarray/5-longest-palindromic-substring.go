package main

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	res := ""

	for i := 0; i < len(s); i++ {
		maxPalindrome(i, i, s, &res)
		maxPalindrome(i, i+1, s, &res)
	}
	return res
}

func maxPalindrome(i, j int, s string, res *string) {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(sub) > len(*res) {
		*res = sub
	}
}
