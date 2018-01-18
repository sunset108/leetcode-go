func findTargetSumWays(nums []int, S int) int {
	if S > 1000 || S < -1000 {
		return 0
	}
	res := make([]int, 2001)
	res[nums[0] + 1000] = 1
	res[-nums[0] + 1000] += 1
	for i := 1; i < len(nums); i++ {
		next := make([]int, 2001)
		for j := nums[i]; j <= 2000 - nums[i]; j++ {
			if res[j] > 0 {
				next[j - nums[i]] += res[j]
				next[j + nums[i]] += res[j]
			}
		}
		res = next
	}
	return res[S + 1000]
}
