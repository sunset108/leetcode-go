func wordBreak(s string, wordDict []string) bool {
    dict := make(map[string]bool, 0)
    for i := 0; i < len(wordDict); i++ {
        dict[wordDict[i]] = true
    }
    n := len(s)
    dp := make([]bool, n + 1)
    dp[0] = true
    for i := 1; i <= n; i++ {
        for j := 0; j < i; j++ {
            if dp[j] && dict[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[n]
}