package main

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < n; i++ {
		len1 := expandCenterAround(s, i, i)
		len2 := expandCenterAround(s, i, i+1)
		len := max(len1, len2)
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandCenterAround(s string, i, j int) int {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return j - i - 1
}
