func isValid(s string) bool {
    stack := make([]byte, 0)
    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            stack = append(stack, s[i])
        } else {
            n := len(stack)
            if n == 0 || (stack[n - 1] == '(' && s[i] != ')') || (stack[n - 1] == '{' && s[i] != '}') || (stack[n - 1] == '[' && s[i] != ']') {
                return false
            }
            stack = stack[:n-1]
        }
    }
    return len(stack) == 0
}