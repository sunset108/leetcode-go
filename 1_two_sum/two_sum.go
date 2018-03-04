func twoSum(nums []int, target int) []int {
	valMap := make(map[int]int, 0)
	for i, v := range nums {
		if _, ok := valMap[target-v]; ok {
			return []int{valMap[target-v], i}
		}
		valMap[v] = i
	}
	return []int{-1, -1}
}
