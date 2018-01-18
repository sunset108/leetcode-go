func findSubstringInWraproundString(p string) int {
	bytes := []byte(p)
	counts := make([]int, 26)
	var curLen, res int
	for i, b := range bytes {
		if i > 0 && (bytes[i - 1] + 1 == bytes[i] || bytes[i - 1] == bytes[i] + 25) {
			curLen++
		} else {
			curLen = 1
		}
		cur := b - 'a'
			counts[cur] = int(math.Max(float64(counts[cur]), float64(curLen)))
	}
	for _, v := range counts {
		res += v
	}
	return res
}
