func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
    memo := make(map[string]bool, 0)
	n := len(nums)
	if n < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < n - 2; i++ {
		j, k := i + 1, n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
        key := fmt.Sprintf("%d%d%d", nums[i], nums[j], nums[k])
				if _, ok := memo[key]; !ok {
					res = append(res, []int{nums[i], nums[j], nums[k]})
					memo[key] = true
				}
        k--
				j++
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return res
}