func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n + 1)
	}
    dp[0][0] = true
	s1, p1 := []byte(s), []byte(p)
    for j := 2; j <= n; j++ {
        dp[0][j] = p1[j - 1] == '*' && dp[0][j - 2]
    }
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
            if j > 1 && p1[j - 1] == '*' {
                dp[i][j] = dp[i][j - 2] || ((s1[i - 1] == p1[j - 2] || p1[j - 2] == '.') && dp[i - 1][j])
            } else {
                dp[i][j] = (s1[i - 1] == p1[j - 1] || p1[j - 1] == '.') && dp[i - 1][j - 1]
            }
		}
	}
	return dp[m][n]
}