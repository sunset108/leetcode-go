func largestRectangleArea(heights []int) int {
    stack := make([]int, 0)
    var max, i, minI int
    for i <= len(heights) {
        n := len(stack)
        if n == 0 || (i < len(heights) && heights[stack[n - 1]] <= heights[i]) {
            stack = append(stack, i)
            i++
        } else {
            preI := stack[n - 1]    
            if n == 1 {
                minI = 0
            } else {
                minI = stack[n - 2] + 1
            }
            if max < (i - minI) * heights[preI] {
                max = (i - minI) * heights[preI]
            }
            stack = stack[:n - 1]
        }
    }
    return max
}