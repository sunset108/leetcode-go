func firstMissingPositive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 1
	}
	i := 0
	for i < n {
		tmp := nums[i]
		if tmp > 0 && tmp <= n && tmp != n + 1 && tmp != nums[tmp-1] {
			nums[i], nums[tmp-1] = nums[tmp-1], nums[i]
		} else {
			i++
		}
		
	}
	for i = 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return i + 1
}