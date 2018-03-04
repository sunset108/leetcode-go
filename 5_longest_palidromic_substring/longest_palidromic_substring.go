func longestPalindrome(s string) string {
	n := len(s)
	bytes := []byte(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	min, max := 0, 0
	for j := 0; j < n; j++ {
		for i := j; i >= 0; i-- {
			if bytes[i] == bytes[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i > max-min {
					min, max = i, j
				}
			}
		}
	}
	return s[min : max+1]
}