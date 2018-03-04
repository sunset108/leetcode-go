func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 1 {
		return find(nums1, nums2, length/2+1)
	} else {
		return (find(nums1, nums2, length/2) + find(nums1, nums2, length/2+1)) / 2
	}
}

func find(nums1 []int, nums2 []int, k int) float64 {
	if len(nums1) > len(nums2) {
		return find(nums2, nums1, k)
	}
	if len(nums1) == 0 {
		return float64(nums2[k-1])
	}
	if k == 1 {
		return math.Min(float64(nums1[0]), float64(nums2[0]))
	}
	p1 := int(math.Min(float64(k/2), float64(len(nums1))))
	p2 := k - p1
	if nums1[p1-1] < nums2[p2-1] {
		return find(nums1[p1:], nums2, p2)
	}
	if nums1[p1-1] > nums2[p2-1] {
		return find(nums1, nums2[p2:], p1)
	}
	return float64(nums1[p1-1])
}