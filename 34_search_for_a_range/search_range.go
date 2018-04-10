func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    return search(nums, 0, len(nums) - 1, target)
}

func search(nums []int, left, right, target int) []int {
    if left == right {
        if nums[left] == target {
            return []int{left, left}
        }
    } else {
        mid := left + (right - left) / 2
        l := search(nums, left, mid, target)
        r := search(nums, mid + 1, right, target)
        if l[0] != -1 && r[0] != -1 && l[1] + 1 == r[0] {
            return []int{l[0], r[1]}
        } else if l[0] == -1 {
            return r
        } else if r[0] == -1 {
            return l
        }
    }
    return []int{-1, -1}
}