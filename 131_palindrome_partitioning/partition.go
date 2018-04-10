var res [][]string

func partition(s string) [][]string {
    res = make([][]string, 0)
    dfs([]string{}, s)
    return res
}

func dfs(list []string, s string) {
    if len(s) == 0 {
        list1 := make([]string, len(list))
        copy(list1, list)
        res = append(res, list1)
        return
    }
    for i := 1; i <= len(s); i++ {
        s1 := s[:i]
        if isPalindrome(s1) {
            dfs(append(list, s1), s[i:])
        }
    }
}

func isPalindrome(s string) bool {
    i, j := 0, len(s) - 1
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}