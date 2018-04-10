func convert(s string, numRows int) string {
    n := len(s)
    if n == 0 || numRows <= 1 {
        return s
    }
    res := ""
    for i := 0; i < numRows; i++ {
        for j := i; j < n; j += (numRows - 1) * 2 {
            res += string(s[j])
            if i != 0 && i != numRows - 1 && j + (numRows - i - 1) * 2 < n {
                res += string(s[j + (numRows - i - 1) * 2])
            }
        }
    }
    return res
}