package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}

	ans, cur := float64(1), x
	for n > 0 {
		if (n & 1) == 1 {
			ans *= cur
		}
		cur *= cur
		n /= 2
	}
	return ans
}

//func main() {
//	x, n := 2.00000, 10
//	fmt.Printf("%v \n", myPow(x, n))
//}
