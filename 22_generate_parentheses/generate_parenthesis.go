var res []string

func generateParenthesis(n int) []string {
    res = make([]string, 0)
    if n <= 0 {
        return res
    }
    dfs(n, n, []byte{})
    return res
}

func dfs(n1, n2 int, bytes []byte) {
    if n1 == 0 && n2 == 0 {
        res = append(res, string(bytes))
        return
    }
    if n1 > 0 {
        dfs(n1 - 1, n2, append(bytes, '('))
    }
    if n2 > n1 {
        dfs(n1, n2 - 1, append(bytes, ')'))
    }
}