func monotoneIncreasingDigits(N int) int {
	str := []byte(strconv.Itoa(N))
	i := 1
	for ; i < len(str) && str[i - 1] <= str[i]; i++ {}
	for ; i > 0 && i < len(str) && str[i - 1] > str[i] ; i-- {
		str[i - 1]--
	}
	for i++; i < len(str); i++ {
		str[i] = '9'
	}
	res, _ := strconv.Atoi(string(str))
	return res
}
