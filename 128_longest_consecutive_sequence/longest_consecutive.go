func longestConsecutive(nums []int) int {
    cache := make(map[int]int, 0)
    longest := 0
    for _, num := range(nums) {
        if _, ok := cache[num]; !ok {
            left := cache[num - 1]
            right := cache[num + 1]
            sum := left + 1 + right
            cache[num] = sum
            cache[num - left] = sum
            cache[num + right] = sum
            if longest < sum {
                longest = sum
            }
        }
    }
    return longest
}