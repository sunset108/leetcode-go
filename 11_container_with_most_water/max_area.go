func maxArea(height []int) int {
    left, right, max, tmp := 0, len(height) - 1, 0, 0
    for left < right {
        if height[left] <= height[right] {
            tmp = height[left] * (right - left)
            left++
        } else {
            tmp = height[right] * (right - left)
            right--
        }
        if max < tmp {
            max = tmp
        }
    }
    return max
}