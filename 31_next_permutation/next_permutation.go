func nextPermutation(nums []int)  {
    n := len(nums)
    if n <= 1 {
        return
    }
    i := n - 2
    for ; i >= 0; i-- {
        if nums[i] < nums[i + 1] {
            j := n - 1
            for ; j > i; j-- {
                if nums[i] < nums[j] {
                    break
                }
            }
            nums[i] = nums[i] ^ nums[j]
            nums[j] = nums[i] ^ nums[j]
            nums[i] = nums[i] ^ nums[j]
            sort.Ints(nums[i + 1:])
            return
        }
    }
    sort.Ints(nums)
    return
}