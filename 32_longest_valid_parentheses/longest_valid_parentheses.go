func longestValidParentheses(s string) int {
    n := len(s)
    if n == 0 {
        return 0
    }
    dp := make([]int, n)
    var max int
    for i := n - 2; i >= 0; i-- {
        if s[i] == '(' {
            j := i + dp[i + 1] + 1
            if j < n && s[j] == ')' {
                dp[i] = j - i + 1
                if j + 1 < n {
                    dp[i] += dp[j + 1]
                }
            }
            if max < dp[i] {
                max = dp[i]
            }
        }
    }
    return max
}