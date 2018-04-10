var res [][]int

func subsets(nums []int) [][]int {
    sort.Ints(nums)
    res = make([][]int, 0)
    dfs(nums, []int{})
    return res
}

func dfs(nums []int, r []int) {
    t := make([]int, len(r))
    copy(t, r)
    res = append(res, t)
    for i := 0; i < len(nums); i++ {
        dfs(nums[i + 1:], append(r, nums[i]))
    }
}