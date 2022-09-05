package main

func removeDuplicateLetters(s string) string {
	stk, freq, instk := make([]byte, 0), make(map[byte]int), make([]bool, 256)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		freq[c]--

		if instk[c] {
			continue
		}
		for len(stk) > 0 && stk[len(stk)-1] > c {
			if freq[stk[len(stk)-1]] == 0 {
				break
			}

			instk[stk[len(stk)-1]] = false
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, c)
		instk[c] = true
	}
	return string(stk)
}
