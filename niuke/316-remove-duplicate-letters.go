package main

func removeDuplicateLetters(s string) string {
	m, stack, instak := make(map[byte]int), make([]byte, 0), make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		m[c]--
		if instak[c] == true {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > c && m[stack[len(stack)-1]] > 0 {
			d := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			instak[d] = false
		}
		stack = append(stack, c)
		instak[c] = true
	}
	return string(stack)
}
