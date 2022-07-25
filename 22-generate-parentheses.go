package main

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := make([]string, 0)
	backtrackParenth(n, n, "", &res)
	return res
}

func backtrackParenth(left, right int, tmp string, res *[]string) {
	if right < left {
		return
	}
	if left < 0 || right < 0 {
		return
	}
	if left == 0 && right == 0 {
		*res = append(*res, tmp)
		return
	}

	tmp = tmp + "("
	backtrackParenth(left-1, right, tmp, res)
	tmp = tmp[:len(tmp)-1]

	tmp = tmp + ")"
	backtrackParenth(left, right-1, tmp, res)
	tmp = tmp[:len(tmp)-1]
}

//func main() {
//	n := 3
//	fmt.Printf("%v \n", generateParenthesis(n))
//}
