package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	for _, c := range s {
		m[byte(c)]++
	}

	for _, c := range t {
		if v, ok := m[byte(c)]; ok {
			if v > 0 {
				m[byte(c)]--
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
