package main

func minWindow(s string, t string) string {
	if s == "" {
		return ""
	}
	if t == "" {
		return ""
	}
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, begin, length, count := 0, 0, len(s)+1, 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++
		if _, ok := need[c]; ok {
			if need[c] == window[c] {
				count++
			}
		}

		for count == len(need) {
			if right-left+1 < length {
				length = right - left + 1
				begin = left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					count--
				}
			}
			window[d]--
		}
	}
	if length == len(s)+1 {
		return ""
	}
	return s[begin : begin+length]
}

//func main() {
//	s, t := "ADOBECODEBANC", "ABC"
//	fmt.Printf("%v \n", minWindow(s, t))
//}
