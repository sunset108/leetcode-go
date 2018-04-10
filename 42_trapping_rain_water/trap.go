func trap(height []int) int {
    if len(height) <= 2 {
        return 0
    }
    left, right := 0, len(height) - 1
    var res int
    for left < right {
        if height[left] <= height[right] {
            min := height[left]
            left++
            for left < right && height[left] < min {
                res += min - height[left]
                left++
            }
        } else {
            min := height[right]
            right--
            for left < right && height[right] < min {
                res += min - height[right]
                right--
            }
        }
    }
    return res
}