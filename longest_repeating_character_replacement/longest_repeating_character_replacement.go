func characterReplacement(s string, k int) int {
    var start, maxNum, res int
    nums := make(map[int]int)
    for end := 0; end < len(s); end++ {
        i := int(s[end] - 'A')
        nums[i]++
        if nums[i] > maxNum {
            maxNum = nums[i]
        }
        if end - start + 1 - maxNum > k {
            nums[int(s[start] - 'A')]--
            start++
        }
        if res < end - start + 1 {
            res = end - start + 1
        }
    }
    return res
}