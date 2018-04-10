func isPalindrome(s string) bool {
    if len(s) == 0 {
        return true
    }
    bytes := []byte(s)
    i, j := 0, len(s) - 1
    for i < j {
        if !isValid(bytes[i]) {
            i++
        } else if !isValid(bytes[j]) {
            j--
        } else {
            if (int(bytes[i]) + 32 - 'a') % 32 != (int(bytes[j]) + 32 - 'a') % 32 {
                return false
            }
            i++
            j--
        }
    }
    return true
}

func isValid(b byte) bool {
    if (b >= '0' && b <= '9') || (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
        return true
    }
    return false
}