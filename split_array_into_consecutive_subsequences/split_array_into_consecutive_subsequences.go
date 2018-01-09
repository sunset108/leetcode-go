func isPossible(nums []int) bool {
	feq, need := make(map[int]int), make(map[int]int)
	for _, num := range nums {
		feq[num]++
	}
	for _, num := range nums {
		if feq[num] != 0 {
			if need[num] > 0 {
				need[num]--
				need[num+1]++
			} else if feq[num+1] > 0 && feq[num+2] > 0 {
				feq[num+1]--
				feq[num+2]--
				need[num+3]++
			} else {
				return false
			}
			feq[num]--
		}
	}
	return true
}
