func sortColors(nums []int)  {
    n := len(nums)
    if n <= 1 {
        return
    }
    i ,s, e := 0, 0, n - 1
    for i <= e {
        switch nums[i] {
        case 0:
            nums[i], nums[s] = nums[s], nums[i]
            s++
            i++
        case 1:
            i++
        case 2:
            nums[i], nums[e] = nums[e], nums[i]
            e--
        }
    }
}