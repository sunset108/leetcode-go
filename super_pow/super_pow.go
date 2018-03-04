func superPow(a int, b []int) int {
	res := 1
	for i := len(b) - 1; i >= 0; i-- {
		res = res * quickPow(a, b[i]) % 1337
		a = quickPow(a, 10)
	}
	return res
}

func quickPow(a int, b int) int {
	a = a % 1337
	res := 1
	for ; b > 0; b >>= 1 {
		if b & 1 == 1 {
			res = res * a % 1337
		}
		a = a * a % 1337
	}
	return res
}
