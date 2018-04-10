func minWindow(s string, t string) string {
    if len(s) < len(t) {
        return ""
    }
    var sCount, tCount [256]int
    for i := 0; i < len(t); i++ {
        tCount[t[i]]++
    }
    minLen, count := len(s) + 1, 0
    var minStr string
    for l, r := 0, 0; r < len(s); r++ {
        c := s[r]
        if tCount[c] != 0 {
            sCount[c]++
            if sCount[c] <= tCount[c] {
                count++
            }
            for count == len(t) {
                if minLen > r - l + 1{
                    minLen = r - l + 1
                    minStr = s[l:r + 1]
                }
                c1 := s[l]
                if sCount[c1] != 0 {
                    sCount[c1]--
                    if sCount[c1] < tCount[c1] {
                        count--
                    }
                }
                l++
            }
        }
    }
    return minStr
}