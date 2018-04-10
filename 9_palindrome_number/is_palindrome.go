func isPalindrome(x int) bool {
    res, y := x, 0
    for res > 0 {
        y = y * 10 + (res % 10)
        res /= 10
    }
    return x == y
}