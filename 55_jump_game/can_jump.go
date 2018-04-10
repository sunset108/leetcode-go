func canJump(nums []int) bool {
    n := len(nums)
    if n == 1 {
        return true
    }
    max := 0
    for i := 0; i < n - 1; i++ {
        if max < i + nums[i] {
            max = i + nums[i]
        }
        if max < i + 1 {
            return false
        }
        if max >= n - 1 {
            return true
        }
    }
    return false
}