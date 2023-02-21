package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2, res := []byte(num1), []byte(num2), make([]int, len(num1)+len(num2))
	for i := 0; i < len(n1); i++ {
		for j := 0; j < len(n2); j++ {
			res[i+j+1] += int(n1[i]-'0') * int(n2[j]-'0')
		}
	}
	for i := len(res) - 1; i >= 1; i-- {
		res[i-1] += res[i] / 10
		res[i] %= 10
	}
	if res[0] == 0 {
		res = res[1:]
	}
	tmp := make([]byte, len(res))
	for i := 0; i < len(res); i++ {
		tmp[i] = '0' + byte(res[i])
	}
	return string(tmp)
}
