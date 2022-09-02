package main

func minWindow(s string, t string) string {
	left, right, freq, window, count, res, l := 0, 0, make(map[byte]int), make(map[byte]int), 0, "", len(s)+1
	for i := 0; i < len(t); i++ {
		freq[t[i]]++
	}

	for right < len(s) {
		c := s[right]
		right++

		if v, ok := freq[c]; ok {
			window[c]++
			if window[c] == v {
				count++
			}
		}
		for count == len(freq) {
			if l > right-left {
				l = right - left
				res = s[left:right]
			}
			d := s[left]
			left++

			if v, ok := freq[d]; ok {
				if window[d] == v {
					count--
				}
				window[d]--
			}
		}
	}
	return res
}

//
//func main() {
//	s, t := "ab", "a"
//	fmt.Printf("%v \n", minWindow(s, t))
//}
