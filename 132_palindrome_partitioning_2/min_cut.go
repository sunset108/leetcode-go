func minCut(s string) int {
    n := len(s)
    if n <= 1 {
        return 0
    }
    dp := make([][]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]bool, n)
        dp[i][i] = true
    }
    count := make([]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        count[i] = n - i
        for j := i; j < n; j++ {
            if s[i] == s[j] {
                if j - i <= 2 || dp[i + 1][j - 1] {
                    dp[i][j] = true
                    if count[i] > count[j + 1] + 1 {
                        count[i] = count[j + 1] + 1
                    }
                }
            }
        }
    }
    return count[0] - 1
}