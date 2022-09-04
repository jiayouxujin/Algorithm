package singlestack

func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}
	stk := make([]byte, 0)
	for i := 0; i < len(num); i++ {
		c := num[i]
		for k > 0 && len(stk) > 0 && stk[len(stk)-1] > c {
			stk = stk[:len(stk)-1]
			k--
		}
		stk = append(stk, c)
	}
	stk = stk[:len(stk)-k]
	for len(stk) > 1 && stk[0] == '0' {
		stk = stk[1:]
	}
	return string(stk)
}
